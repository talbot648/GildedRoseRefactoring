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
func (item *Item) UpdateQuality() {
	if item.Name != "Backstage passes to a TAFKAL80ETC concert" {
		if item.Name != "Sulfuras, Hand of Ragnaros" {
			item.decreaseQualityByOne()
		}
		if isConjured(item) {
			item.decreaseQualityByOne()
		}
	} else {
		item.increaseQualityByOne()
		if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
			if item.SellIn < 11 {
				item.increaseQualityByOne()
			}
			if item.SellIn < 6 {
				{
					item.increaseQualityByOne()
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
				{
					if item.Name != "Sulfuras, Hand of Ragnaros" {
						item.decreaseQualityByOne()
					}
				}
			} else {
				item.Quality = 0
			}
		} else {
			{
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
	if isItemQualityAboveZero(item) {
		item.Quality--
	}
}

func (item *Item) decreaseSellInByOne() {
	item.SellIn--
}

func (item *Item) increaseQualityByOne() {
	if isItemQualityBelowFifty(item) {
		item.Quality++
	}
}
