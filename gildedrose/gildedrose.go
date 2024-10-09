package gildedrose

import "strings"

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateItems(items []*Item) {
	for _, item := range items {
		switch item.Name {
		case "Aged Brie":
			item.UpdateAgedBrie()
		case "Backstage passes to a TAFKAL80ETC concert":
			item.UpdateBackstagePass()
		case "Sulfuras, Hand of Ragnaros":
		default:
			item.UpdateQuality()
		}
	}
}

func (item *Item) UpdateAgedBrie() {
	item.decreaseSellInByOne()
	item.increaseQualityByOne()

	if isItemSellInDatePassed(item) {
		item.increaseQualityByOne()
	}
}

func (item *Item) UpdateBackstagePass() {
	item.decreaseSellInByOne()

	item.increaseQualityByOne()
	if item.SellIn < 10 {
		item.increaseQualityByOne()
	}
	if item.SellIn < 5 {
		item.increaseQualityByOne()
	}

	if isItemSellInDatePassed(item) {
		item.Quality = 0
	}
}

func (item *Item) UpdateQuality() {

	item.decreaseQualityByOne()

	if isConjured(item) {
		item.decreaseQualityByOne()
	}
	item.decreaseSellInByOne()
	if isItemSellInDatePassed(item) {
		item.decreaseQualityByOne()
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
