package constants

const (
	KUFAR_PARSE_URL   = "https://api.kufar.by/search-api/v2/search/rendered-paginated?cat=1010&cur=USD&gtsy=country-belarus~province-minsk~locality-minsk&lang=ru&prc=r%3A500%2C700&rnt=1&size=30&typ=let"
	REALT_PARSE_URL   = "https://realt.by/_next/data/y2lP1_Q8X4RnFpI5nfo8e/rent/flat-for-long.json?sortType=createdAt&page=1"
	ONLINER_PARSE_URL = "https://r.onliner.by/sdapi/ak.api/search/apartments?price%5Bmin%5D=500&price%5Bmax%5D=700&currency=usd&bounds%5Blb%5D%5Blat%5D=53.820922446131306&bounds%5Blb%5D%5Blong%5D=27.30583190917969&bounds%5Brt%5D%5Blat%5D=53.97466657389324&bounds%5Brt%5D%5Blong%5D=27.818069458007816&order=created_at%3Adesc&page=1&v=0.6848485053787079"
)

const (
	T_KUFAR   = 1
	T_REALT   = 2
	T_ONLINER = 3
)

const KUFAR_GALLERY_PATH = "https://rms6.kufar.by/v1/gallery/%s"

const (
	LOWER_SEARCH_PRICE  = 500
	HIGHER_SEARCH_PRICE = 700
)
