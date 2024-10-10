package gildedrose

type AgedBrie struct {
	*Item
}

type BackstagePass struct {
	*Item
}

type Conjured struct {
	*Item
}

type RegularItem struct {
	*Item
}

type Sulfuras struct {
	*Item
}

type Item struct {
	Name            string
	SellIn, Quality int
}
