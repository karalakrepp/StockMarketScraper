# StockMarketScraper

## Kullanım
1. Projeyi bilgisayarınıza kopyalayın:

   ```bash
    git clone https://github.com/kullanici/adiniz/stock-scraper.git


2. Gerekli bağımlılıkları yükleyin:

    ```bash
    go get -u github.com/gocolly/colly/v2


3. Programı çalıştırın:

    ```bash
    go run main.go


## Özellikler
Program, belirli bir liste üzerindeki hisse senetlerinin şirket adı, hisse fiyatı ve yüzde değişim bilgilerini toplar.
Toplanan verileri stocks.csv adlı bir CSV dosyasına kaydeder.