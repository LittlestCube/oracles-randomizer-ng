package hints

// rules:
// - names of specific items get capitalized
// - articles *usually* appear as in the "you got x" text, but are always a/an
//   in the case of nonspecific items ("a sword")

type itemName struct {
	article string
	name    string
	plural  bool
}

var itemMap = map[string]itemName{
	"wooden shield":      {"a", "shield", false},
	"iron shield":        {"a", "shield", false},
	"bombs, 10":          {"", "Bombs", false},
	"sword 1":            {"a", "sword", false},
	"sword 2":            {"a", "sword", false},
	"ricky's flute":      {"", "Ricky's Flute", false},
	"dimitri's flute":    {"", "Dimitri's Flute", false},
	"moosh's flute":      {"", "Moosh's Flute", false},
	"shovel":             {"the", "Shovel", false},
	"feather 1":          {"a", "feather", false},
	"feather 2":          {"a", "feather", false},
	"satchel 1":          {"a", "Seed Satchel", false},
	"satchel 2":          {"a", "Seed Satchel", false},
	"ember tree seeds":   {"", "Ember Seeds", true},
	"scent tree seeds":   {"", "Scent Seeds", true},
	"pegasus tree seeds": {"", "Pegasus Seeds", true},
	"gale tree seeds":    {"", "Gale Seeds", true},
	"mystery tree seeds": {"", "Mystery Seeds", true},
	"heart container":    {"a", "Heart Container", false},
	"piece of heart":     {"a", "Piece of Heart", false},
	"gasha seed":         {"a", "Gasha Seed", false},
	"ricky's gloves":     {"", "Ricky's Gloves", true},
	"bomb flower":        {"a", "Bomb Flower", false},

	// seasons-specific
	"boomerang 1":      {"a", "boomerang", false},
	"boomerang 2":      {"a", "boomerang", false},
	"spring":           {"the", "Rod of Spring", false},
	"summer":           {"the", "Rod of Summer", false},
	"autumn":           {"the", "Rod of Autumn", false},
	"winter":           {"the", "Rod of Winter", false},
	"magnet gloves":    {"the", "Magnetic Gloves", true},
	"slingshot 1":      {"a", "slingshot", false},
	"slingshot 2":      {"a", "slingshot", false},
	"bracelet":         {"the", "Power Bracelet", false},
	"fool's ore":       {"", "Fool's Ore", false},
	"rare peach stone": {"a", "Piece of Heart", false},
	"flippers":         {"", "Zora's Flippers", true},
	"gnarled key":      {"the", "Gnarled Key", false},
	"floodgate key":    {"the", "Floodgate Key", false},
	"dragon key":       {"the", "Dragon Key", false},
	"star ore":         {"a", "Star-Shaped Ore", false},
	"ribbon":           {"a", "Ribbon", false},
	"spring banana":    {"a", "Spring Banana", false},
	"rusty bell":       {"the", "Rusty Bell", false},
	"treasure map":     {"the", "Treasure Map", false},
	"round jewel":      {"the", "Round Jewel", false},
	"pyramid jewel":    {"the", "Pyramid Jewel", false},
	"square jewel":     {"the", "Square Jewel", false},
	"x-shaped jewel":   {"the", "X-Shaped Jewel", false},
	"red ore":          {"", "Red Ore", false},
	"blue ore":         {"", "Blue Ore", false},
	"hard ore":         {"", "Hard Ore", false},
	"member's card":    {"a", "Member's Card", false},
	"master's plaque":  {"the", "Master's Plaque", false},

	// ages-specific
	"cane":            {"the", "Cane of Somaria", false},
	"boomerang":       {"a", "Boomerang", false},
	"switch hook 1":   {"a", "hook", false},
	"switch hook 2":   {"a", "hook", false},
	"seed shooter":    {"the", "Seed Shooter", false},
	"harp 1":          {"a", "tune", false},
	"harp 2":          {"a", "tune", false},
	"harp 3":          {"a", "tune", false},
	"bracelet 1":      {"a", "bracelet", false},
	"bracelet 2":      {"a", "bracelet", false},
	"feather":         {"", "Roc's Feather", false},
	"flippers 1":      {"", "flippers", true},
	"flippers 2":      {"", "flippers", true},
	"graveyard key":   {"the", "Graveyard Key", false},
	"crown key":       {"the", "Crown Key", false},
	"old mermaid key": {"the", "Old Mermaid Key", false},
	"mermaid key":     {"the", "Mermaid Key", false},
	"library key":     {"the", "Library Key", false},
	"tuni nut":        {"the", "Tuni Nut", false},
	"scent seedling":  {"a", "Seedling", false},
	"zora scale":      {"the", "Zora Scale", false},
	"tokay eyeball":   {"the", "Tokay Eyeball", false},
	"fairy powder":    {"", "Fairy Powder", false},
	"cheval rope":     {"", "Cheval Rope", false},
	"island chart":    {"an", "Island Chart", false},
	"book of seals":   {"the", "Book of Seals", false},
	"goron letter":    {"a", "Letter of Introduction", false},
	"lava juice":      {"", "Lava Juice", false},
	"brother emblem":  {"the", "Brother Emblem", false},
	"goron vase":      {"the", "Goron Vase", false},
	"goronade":        {"", "Goronade", false},
	"rock brisket":    {"", "Rock Brisket", false},

	// rings
	"friendship ring": {"the", "Friendship Ring", false},
	"power ring L-1":  {"the", "Power Ring L-1", false},
	"power ring L-2":  {"the", "Power Ring L-2", false},
	"power ring L-3":  {"the", "Power Ring L-3", false},
	"armor ring L-1":  {"the", "Armor Ring L-1", false},
	"armor ring L-2":  {"the", "Armor Ring L-2", false},
	"armor ring L-3":  {"the", "Armor Ring L-3", false},
	"red ring":        {"the", "Red Ring", false},
	"blue ring":       {"the", "Blue Ring", false},
	"green ring":      {"the", "Green Ring", false},
	"cursed ring":     {"the", "Cursed Ring", false},
	"expert's ring":   {"the", "Expert's Ring", false},
	"blast ring":      {"the", "Blast Ring", false},
	"rang ring L-1":   {"the", "Rang Ring L-1", false},
	"GBA time ring":   {"the", "GBA Time Ring", false},
	"maple's ring":    {"the", "Maple's Ring", false},
	"steadfast ring":  {"the", "Steadfast Ring", false},
	"pegasus ring":    {"the", "Pegasus Ring", false},
	"toss ring":       {"the", "Toss Ring", false},
	"heart ring L-1":  {"the", "Heart Ring L-1", false},
	"heart ring L-2":  {"the", "Heart Ring L-2", false},
	"swimmer's ring":  {"the", "Swimmer's Ring", false},
	"charge ring":     {"the", "Charge Ring", false},
	"light ring L-1":  {"the", "Light Ring L-1", false},
	"light ring L-2":  {"the", "Light Ring L-2", false},
	"bomber's ring":   {"the", "Bomber's Ring", false},
	"green luck ring": {"the", "Green Luck Ring", false},
	"blue luck ring":  {"the", "Blue Luck Ring", false},
	"gold luck ring":  {"the", "Gold Luck Ring", false},
	"red luck ring":   {"the", "Red Luck Ring", false},
	"green holy ring": {"the", "Green Holy Ring", false},
	"blue holy ring":  {"the", "Blue Holy Ring", false},
	"red holy ring":   {"the", "Red Holy Ring", false},
	"snowshoe ring":   {"the", "Snowshoe Ring", false},
	"roc's ring":      {"the", "Roc's Ring", false},
	"quicksand ring":  {"the", "Quicksand Ring", false},
	"red joy ring":    {"the", "Red Joy Ring", false},
	"blue joy ring":   {"the", "Blue Joy Ring", false},
	"gold joy ring":   {"the", "Gold Joy Ring", false},
	"green joy ring":  {"the", "Green Joy Ring", false},
	"discovery ring":  {"the", "Discovery Ring", false},
	"rang ring L-2":   {"the", "Rang Ring L-2", false},
	"octo ring":       {"the", "Octo Ring", false},
	"moblin ring":     {"the", "Moblin Ring", false},
	"like-like ring":  {"the", "Like-Like Ring", false},
	"subrosian ring":  {"the", "Subrosian Ring", false},
	"first gen ring":  {"the", "First Gen Ring", false},
	"spin ring":       {"the", "Spin Ring", false},
	"bombproof ring":  {"the", "Bombproof Ring", false},
	"energy ring":     {"the", "Energy Ring", false},
	"dbl. edged ring": {"the", "Dbl. Edged Ring", false},
	"GBA nature ring": {"the", "GBA Nature Ring", false},
	"slayer's ring":   {"the", "Slayer's Ring", false},
	"rupee ring":      {"the", "Rupee Ring", false},
	"victory ring":    {"the", "Victory Ring", false},
	"sign ring":       {"the", "Sign Ring", false},
	"100th ring":      {"the", "100th Ring", false},
	"whisp ring":      {"the", "Whisp Ring", false},
	"gasha ring":      {"the", "Gasha Ring", false},
	"peace ring":      {"the", "Peace Ring", false},
	"zora ring":       {"the", "Zora Ring", false},
	"fist ring":       {"the", "Fist Ring", false},
	"whimsical ring":  {"the", "Whimsical Ring", false},
	"protection ring": {"the", "Protection Ring", false},

	// just call all of these "Rupees" instead of specific amounts. otherwise
	// it might imply a *total* number of rupees in a location.
	"rupees, 1":   {"", "Rupees", true},
	"rupees, 5":   {"", "Rupees", true},
	"rupees, 10":  {"", "Rupees", true},
	"rupees, 20":  {"", "Rupees", true},
	"rupees, 30":  {"", "Rupees", true},
	"rupees, 50":  {"", "Rupees", true},
	"rupees, 100": {"", "Rupees", true},
	"rupees, 200": {"", "Rupees", true},

	// unused
	"rod":           {"the", "Rod of Seasons", false},
	"bombchus":      {"", "Bombchus", true},
	"strange flute": {"a", "Strange Flute", false},
	"ring box L-1":  {"a", "ring box", false},
	"ring box L-2":  {"a", "ring box", false},
	"small key":     {"a", "Small Key", false},
	"d0 small key":  {"a", "Small Key", false},
	"d1 small key":  {"a", "Small Key", false},
	"d2 small key":  {"a", "Small Key", false},
	"d3 small key":  {"a", "Small Key", false},
	"d4 small key":  {"a", "Small Key", false},
	"d5 small key":  {"a", "Small Key", false},
	"d6 small key":  {"a", "Small Key", false},
	"d7 small key":  {"a", "Small Key", false},
	"d8 small key":  {"a", "Small Key", false},
	"boss key":      {"a", "Boss Key", false},
	"d1 boss key":   {"a", "Boss Key", false},
	"d2 boss key":   {"a", "Boss Key", false},
	"d3 boss key":   {"a", "Boss Key", false},
	"d4 boss key":   {"a", "Boss Key", false},
	"d5 boss key":   {"a", "Boss Key", false},
	"d6 boss key":   {"a", "Boss Key", false},
	"d7 boss key":   {"a", "Boss Key", false},
	"d8 boss key":   {"a", "Boss Key", false},
	"compass":       {"a", "Compass", false},
	"dungeon map":   {"a", "Dungeon Map", false},
	"slate 1":       {"a", "Slate", false},
	"slate 2":       {"a", "Slate", false},
	"slate 3":       {"a", "Slate", false},
	"slate 4":       {"a", "Slate", false},
}
