package gildedrose

import "strings"

type Update interface {
	update()
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
