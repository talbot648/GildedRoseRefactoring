package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func TestQualityDecreasesByOneAfterEachDay(t *testing.T) {
	item := givenTestItem("test", 7, 5)

	gildedrose.UpdateItems(item)
	got := item[0].Quality
	want := 4

	if got != want {
		t.Errorf("error when trying to decrease quality, got %v, expected %v", got, want)
	}
}

func TestQualityDegradeSpeedDoublesAfterPassedSellIn(t *testing.T) {
	item := givenTestItem("test", 0, 10)

	gildedrose.UpdateItems(item)
	got := item[0].Quality
	want := 8

	if got != want {
		t.Errorf("error when trying to double decreasing speed, got %v, expected %v", got, want)
	}
}

func TestQualityCannotGoBelowZero(t *testing.T) {
	item := givenTestItem("test", 2, 0)

	gildedrose.UpdateItems(item)
	got := item[0].Quality
	want := 0

	if got != want {
		t.Errorf("error quality went below zero, got %v, expected %v", got, want)
	}
}

func TestAgedBrieQualityIncreasesAfterUpdate(t *testing.T) {
	item := givenTestItem("Aged Brie", 2, 5)

	gildedrose.UpdateItems(item)
	got := item[0].Quality
	want := 6

	if got != want {
		t.Errorf("error when trying to increase quality, got %v, expected %v", got, want)
	}
}

func TestQualityCannotGoAboveFifty(t *testing.T) {
	item := givenTestItem("Aged Brie", 2, 50)

	gildedrose.UpdateItems(item)
	got := item[0].Quality
	want := 50

	if got != want {
		t.Errorf("error: quality went above fifty, got %v, expected %v", got, want)
	}
}

func TestSellInDecreasesByOneAfterEachDay(t *testing.T) {
	item := givenTestItem("test", 5, 7)

	gildedrose.UpdateItems(item)
	got := item[0].SellIn
	want := 4

	if got != want {
		t.Errorf("error when trying to decrease sellIn, got %v, expected %v", got, want)
	}
}

func TestSulfurasDoesNotReduceInQuality(t *testing.T) {
	item := givenTestItem("Sulfuras, Hand of Ragnaros", 3, 80)

	gildedrose.UpdateItems(item)
	got := item[0].Quality
	want := 80

	if got != want {
		t.Errorf("error when trying to decrease quality, got %v, expected %v", got, want)
	}
}

func TestSulfurasSellInDoesNotAffectQuality(t *testing.T) {
	item := givenTestItem("Sulfuras, Hand of Ragnaros", -2, 80)

	gildedrose.UpdateItems(item)
	got := item[0].Quality
	want := 80

	if got != want {
		t.Errorf("error when trying to decrease quality, got %v, expected %v", got, want)
	}
}

func TestBackstagePassQualityIncreaseByOneWhenSellInOverTenDays(t *testing.T) {
	item := givenTestItem("Backstage passes to a TAFKAL80ETC concert", 13, 5)

	gildedrose.UpdateItems(item)
	got := item[0].Quality
	want := 6

	if got != want {
		t.Errorf("error when updating quality of backstage pass above ten day sellIn, got %v, expected %v", got, want)
	}
}

func TestBackstagePassQualityIncreaseByTwoWhenSellInBetweenFiveAndTenDays(t *testing.T) {
	item := givenTestItem("Backstage passes to a TAFKAL80ETC concert", 7, 5)

	gildedrose.UpdateItems(item)
	got := item[0].Quality
	want := 7

	if got != want {
		t.Errorf("error when updating quality of backstage pass above ten day sellIn, got %v, expected %v", got, want)
	}
}

func TestBackstagePassQualityIncreaseByThreeWhenSellInLessThanFiveDays(t *testing.T) {
	item := givenTestItem("Backstage passes to a TAFKAL80ETC concert", 3, 5)

	gildedrose.UpdateItems(item)
	got := item[0].Quality
	want := 8

	if got != want {
		t.Errorf("error when updating quality of backstage pass above ten day sellIn, got %v, expected %v", got, want)
	}
}

func TestBackstagePassQualityZeroAfterConcert(t *testing.T) {
	item := givenTestItem("Backstage passes to a TAFKAL80ETC concert", 0, 5)

	gildedrose.UpdateItems(item)
	got := item[0].Quality
	want := 0

	if got != want {
		t.Errorf("error when updating quality of backstage pass above ten day sellIn, got %v, expected %v", got, want)
	}
}

func givenTestItem(name string, SellIn int, Quality int) []*gildedrose.Item {
	var items = []*gildedrose.Item{
		{name, SellIn, Quality},
	}
	return items

}
