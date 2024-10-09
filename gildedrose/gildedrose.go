package gildedrose

import "strings"

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateItems(items []*Item) {
	for _, item := range items {
		item.UpdateQuality()
	}
}

func (item *Item) UpdateQuality() {
	if item.Name != "Aged Brie" && item.Name != "Backstage passes to a TAFKAL80ETC concert" {
		if isItemQualityAboveZero(item) {
			if item.Name != "Sulfuras, Hand of Ragnaros" {
				item.decreaseQualityByOne()
			}
			if isConjured(item) && isItemQualityAboveZero(item) {
				item.decreaseQualityByOne()
			}
		}
	} else {
		if isItemQualityBelowFifty(item) {
			item.increaseQualityByOne()
			if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
				if item.SellIn < 11 {
					if item.Quality < 50 {
						item.increaseQualityByOne()

					}
				}
				if item.SellIn < 6 {
					if item.Quality < 50 {
						item.increaseQualityByOne()
					}
				}
			}
		}
	}

	if item.Name != "Sulfuras, Hand of Ragnaros" {
		item.decreaseSellInByOne()
	}

	if isItemSellInDatePassed(item) {
		if item.Name != "Aged Brie" {
			if item.Name != "Backstage passes to a TAFKAL80ETC concert" {
				if isItemQualityAboveZero(item) {
					if item.Name != "Sulfuras, Hand of Ragnaros" {
						item.decreaseQualityByOne()
					}
				}
			} else {
				item.Quality = 0
			}
		} else {
			if isItemQualityBelowFifty(item) {
				item.increaseQualityByOne()
			}
		}
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
	item.Quality--
}

func (item *Item) decreaseSellInByOne() {
	item.SellIn--
}

func (item *Item) increaseQualityByOne() {
	item.Quality++
}
