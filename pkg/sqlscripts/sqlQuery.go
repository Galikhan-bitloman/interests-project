package sqlscripts

const (
	Calc string = `INSERT INTO "postgres".public."moneyTracking" (salary, necessary_spending, savings, entertainment)
			VALUES ($1, $2, $3, $4)
`
	GetAllCalcData = `SELECT salary, necessary_spending, savings, entertainment, created_at 
				  FROM "postgres".public."moneyTracking"
				  ORDER BY created_at desc
 				 `
	InsertNasaData string = `INSERT INTO "postgres"."public"."nasa_images" (url, explanation, hdurl, date, title)
					  VALUES ($1, $2, $3, $4, $5)`
	GetAllNasaImage string = `SELECT url, explanation, hdurl, date, title FROM "postgres".public."nasa_images" order by date desc`
)
