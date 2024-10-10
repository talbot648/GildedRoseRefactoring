package gildedrose

import "strings"

type Item struct {
	Name            string
	SellIn, Quality int
}

type Update interface {
	update()
}

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

func convertItemType(item *Item) Update {

	switch item.Name {

	case "Aged Brie":
		return &AgedBrie{item}

	case "Backstage passes to a TAFKAL80ETC concert":
		return &BackstagePass{item}

	case "Sulfuras, Hand of Ragnaros":
		return &Sulfuras{item}
	default:
		if isConjured(item) {
			return &Conjured{item}
		}
	}
	return &RegularItem{item}

}

func UpdateItems(items []*Item) {
	for _, item := range items {
		updatedItem := convertItemType(item)
		updatedItem.update()
	}
}

func (s *Sulfuras) update() {
	//nothing needed to update a legendary item
}

func (c *Conjured) update() {
	for i := 0; i < 2; i++ {
		if c.Quality > 0 {
			c.decreaseQualityByOne()
		}
	}
	c.decreaseSellInByOne()
}

func (b *AgedBrie) update() {
	b.decreaseSellInByOne()
	b.increaseQualityByOne()

	if isItemSellInDatePassed(b.Item) {
		b.increaseQualityByOne()
	}
}

func (bp *BackstagePass) update() {
	bp.decreaseSellInByOne()

	bp.increaseQualityByOne()
	if bp.SellIn < 10 {
		bp.increaseQualityByOne()
	}
	if bp.SellIn < 5 {
		bp.increaseQualityByOne()
	}

	if isItemSellInDatePassed(bp.Item) {
		bp.Quality = 0
	}
}

func (r *RegularItem) update() {

	r.decreaseQualityByOne()
	r.decreaseSellInByOne()
	if isItemSellInDatePassed(r.Item) {
		r.decreaseQualityByOne()
	}
}

func isConjured(item *Item) bool {
	return strings.Contains(item.Name, "Conjured")
}

func isItemSellInDatePassed(item *Item) bool {
	return item.SellIn < 0
}

func isItemQualityAboveZero(item *Item) bool {
	return item.Quality > 0
}

func isItemQualityBelowFifty(item *Item) bool {
	return item.Quality < 50
}

func (item *Item) decreaseQualityByOne() {
	if isItemQualityAboveZero(item) {
		item.Quality--
	}
}

func (item *Item) decreaseSellInByOne() {
	isItemSellInDatePassed(item)
	item.SellIn--
}

func (item *Item) increaseQualityByOne() {
	if isItemQualityBelowFifty(item) {
		item.Quality++
	}
}
