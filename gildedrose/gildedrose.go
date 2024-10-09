package gildedrose

import "strings"

type Item struct {
	Name            string
	SellIn, Quality int
}

type UpdateItem interface {
	UpdateItem()
}

type AgedBrie struct {
	*Item
}

func updateItemType(item *Item) UpdateItem {
	return &AgedBrie{item}
}

func UpdateItems(items []*Item) {
	for _, item := range items {

		switch item.Name {

		case "Aged Brie":
			updatedItem := updateItemType(item)
			updatedItem.UpdateItem()

		case "Backstage passes to a TAFKAL80ETC concert":
			item.UpdateBackstagePass()

		case "Sulfuras, Hand of Ragnaros":

		default:
			if isConjured(item) {
				item.UpdateConjured()
			} else {
				item.UpdateQuality()
			}
		}
	}
}

func (item *Item) UpdateConjured() {
	for i := 0; i < 2; i++ {
		if item.Quality > 0 {
			item.Quality--
		}
	}
	item.decreaseSellInByOne()
}

func (b *AgedBrie) UpdateItem() {
	b.decreaseSellInByOne()
	b.increaseQualityByOne()

	if isItemSellInDatePassed(b.Item) {
		b.increaseQualityByOne()
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
