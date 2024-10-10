package gildedrose

func (s *Sulfuras) update() {
	//nothing needed to update a legendary item
}

func (c *Conjured) update() {
	c.decreaseQualityByOne()
	c.decreaseQualityByOne()
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
