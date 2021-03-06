package ais

import (
	"bytes"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/kiberlom/moskovskay/internal/models"

	"github.com/PuerkitoBio/goquery"
)

func Parsing(body []byte) ([]models.Apartment, error) {

	nr := bytes.NewReader(body)

	doc, err := goquery.NewDocumentFromReader(nr)
	if err != nil {
		log.Fatal(err)
	}

	var apartments []models.Apartment
	var details []models.DetailCase

	doc.Find(".table-full-width").Each(func(i int, s *goquery.Selection) {
		s.Find("tbody").Each(func(i int, t *goquery.Selection) {

			apartmentOne := models.Apartment{}

			t.Find(".clickable").Each(func(i int, oneZ *goquery.Selection) {

				oneZ.Find("td").Each(func(i int, kv *goquery.Selection) {

					res := strings.TrimSpace(kv.Text())

					switch i {
					case 0:
						r, err := strconv.Atoi(res)
						if err != nil {
							apartmentOne.Number = -1
							break
						}
						apartmentOne.Number = r

					case 1:
						r, err := strconv.Atoi(res)
						if err != nil {
							apartmentOne.NumberRoom = -1
							break
						}
						apartmentOne.NumberRoom = r

					case 2:
						apartmentOne.TimeCreate = res
					case 3:
						apartmentOne.Priority = res
					case 4:
						apartmentOne.Status = res
					case 5:
						apartmentOne.TimeTransfer = res
					case 6:
						apartmentOne.TimeFinish = res
					}
				})
				apartments = append(apartments, apartmentOne)
			})

			t.Find(".none").Each(func(i int, s *goquery.Selection) {

				detailOne := models.DetailCase{}

				s.Find("td").Find("p").Each(func(i int, deT *goquery.Selection) {

					res := strings.TrimSpace(deT.After("strong").Text())

					switch i {
					case 0:
						detailOne.ApplicantAssessment = res
					case 1:
						detailOne.Cause = res
					case 2:
						detailOne.ApplicationType = res
					case 3:
						detailOne.WorkGroup = res
					case 4:
						detailOne.ControlClaim = res
					case 5:
						detailOne.MES = res
					case 6:
						detailOne.ImplementingOrganization = res
					case 7:
						detailOne.Work = res
					}

				})

				details = append(details, detailOne)
			})

		})

	})

	// ???????? ???? ?????????????? ????????????
	if len(apartments) == 0 {
		return nil, nil
	}

	// ???????????????????? ????????????
	for i := range apartments {
		apartments[i].DetailCase = details[i]
	}

	// ?????????????????????? ???? ???????????? ????????????
	sort.SliceStable(apartments, func(i, j int) bool {
		return apartments[i].Number < apartments[j].Number
	})

	return apartments, nil

}
