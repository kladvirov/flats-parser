package constants

const (
	KufarParseUrl   = "https://api.kufar.by/search-api/v2/search/rendered-paginated?cat=1010&cur=USD&gtsy=country-belarus~province-minsk~locality-minsk&lang=ru&prc=r%3A250%2C500&size=30&typ=let"
	RealtParseUrl   = "https://realt.by/bff/graphql"
	OnlinerParseUrl = "https://r.onliner.by/sdapi/ak.api/search/apartments?price%5Bmin%5D=250&price%5Bmax%5D=500&currency=usd&bounds%5Blb%5D%5Blat%5D=53.820922446131306&bounds%5Blb%5D%5Blong%5D=27.30583190917969&bounds%5Brt%5D%5Blat%5D=53.97466657389324&bounds%5Brt%5D%5Blong%5D=27.818069458007816&order=created_at%3Adesc&page=1&v=0.6848485053787079"
)

const (
	TKufar   = 1
	TRealt   = 2
	TOnliner = 3
)

const KufarGalleryPath = "https://rms6.kufar.by/v1/gallery/%s"

const (
	LowerSearchPrice  = 250
	HigherSearchPrice = 500
)

const GqlQuery = `
query searchObjects($data: GetObjectsByAddressInput!) {
  searchObjects(data: $data) {
    body {
      results {
        uuid
        code
        title
        price
        priceCurrency
        areaTotal
        areaLiving
        rooms
        storey
        storeys
        images
        createdAt
        address
        directionName
        location
        headline
        metroStationName
      }
    }
    ...StatusAndErrors
  }
}
fragment StatusAndErrors on INullResponse {
  success
  errors { code title message field }
}
`
