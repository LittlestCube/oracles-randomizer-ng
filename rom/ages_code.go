package rom

import (
	"strings"
)

func newAgesRomBanks() *romBanks {
	asm, err := newAssembler()
	if err != nil {
		panic(err)
	}

	r := romBanks{
		endOfBank: loadBankEnds("ages"),
		assembler: asm,
	}

	// do this before loading asm files, since the sizes of the tables vary
	// with the number of checks.
	r.replaceRaw(Addr{0x06, 0}, "collectModeTable", makeAgesCollectModeTable())
	r.replaceRaw(Addr{0x38, 0}, "roomTreasures", makeRoomTreasureTable())
	r.replaceRaw(Addr{0x3f, 0}, "owlTextOffsets", string(make([]byte, 0x14*2)))

	r.applyAsmFiles(GameAges,
		[]string{
			"/asm/common.yaml",
			"/asm/ages.yaml",
		},
		[]string{
			"/asm/animals.yaml",
			"/asm/cutscenes.yaml",
			"/asm/gfx.yaml",
			"/asm/item_events.yaml",
			"/asm/item_lookup.yaml",
			"/asm/layouts.yaml",
			"/asm/linked.yaml",
			"/asm/misc.yaml",
			"/asm/rings.yaml",
			"/asm/triggers.yaml",
			"/asm/vars.yaml",

			"/asm/text.yaml", // must go last
		})

	return &r
}

// makes ages-specific additions to the collection mode table.
func makeAgesCollectModeTable() string {
	b := new(strings.Builder)
	table := makeCollectModeTable()
	b.WriteString(table[:len(table)-1]) // strip final ff

	// add eatern symmetry city brother
	b.Write([]byte{0x03, 0x6f, collectModes["touch"]})

	// add ricky and dimitri nuun caves
	b.Write([]byte{0x02, 0xec, collectModes["chest"],
		0x05, 0xb8, collectModes["chest"]})

	b.Write([]byte{0xff})
	return b.String()
}
