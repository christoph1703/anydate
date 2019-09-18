package anydate

import (
	"fmt"
	"testing"
)

const (
	testdata = `Namecheap, Inc.
	                               Transactions  4600 East Washington Street. Suite 305,
	                                                        Phoenix, AZ 852004
	                                                        USA
	        www.namecheap.com      Transaction # 51091820  support@namecheap.com
	         
	       Transaction Date                        :     May 02, 2019 01:10:45 PM
	       Transaction Id                          :     51091820
	       Transaction Type                        :     Order
	       Order ID                                :     44476000
	       Payment Method                          :     Paypal
	       Payment Reference ID                    :
	       Charge Amount                           :     $9.06
	       Status                    2019-07-05              :     Success
	         
	
	Hornbach Baumarkt GmbH
Filiale 682
Industriegebiet B 21

VV ahee7al Bad-Fi-gchau
Tel als (Oy 3a bd AE 0-0

Sra PRRCAB: 309

ART/EAN*8591.000041.589
1 Stuck x
OSB-Verlegepl. 12mm

SUMME [1]

GEGEBEN Maestro

-K-U-N-D-E-N-B-E-L-E-G-

Terminal-ID : 65027542
TA-Nr 027675 BNr 0155

Kartenzahl ung
contactless
MAESTRO

EUR oS. >

PAN HHHHHHAHHHERHHHOOGO
Karte 1 gUltig bis #H/##H
EMV-AID AQ000000043060
VU-Nr 201059258
Genehmi gungs-Nr 994006
Datum 20.04.19 11:30 Uhr

xxx Zahlung erfolgt xxx

AS-Proc-Code = 00 075 00
Capt.-Ref.= 0497

BITTE BELEG AUFBEWAHREN

By

HL

4

Fiskale Transaktionsnu 414
Fiskale Kassennummer : 2

Es bediente Sie:
M. Kaiser

Hinweise zum Umgang mit Ihren Daten
und zu Ihren Rechten finden Sie unter
www.hornbach.at - Daten und Recht
>> Datum = Liefer- und Rechnungsdatum <<

0682. 20.04.2019 11:30 0012 000023 004451

 
Christoph Armster
Armster electronics
Adalbert-Stifter-Straße 11/29
1200 Wien
Austria
VAT number: AT U74587823
Details
Google Cloud - GSuite
..............................................................
Invoice
number
3630081488
..............................................................
Invoice
date
Aug 31, 2019
Total in EUR
..............................................................
Billing
ID
7552-2694-2285
1020
..............................................................
Domain
name
armster-electronics.com
€5.20
Summary for Aug 1, 2019 - Aug 31, 2019
Subtotal in EUR €5.20
VAT (0%) €0.00
Total in EUR €5.20
Services subject to the reverse charge - VAT to be accounted for by the
 
05. Jun 2019

18. August 2010

2018.9.23


15 Februar 2014

`
	resultData = `2019-05-02
2019-07-05
2019-04-20
2019-08-31
2019-08-01
2019-08-31
2019-06-05
2010-08-18
2018-09-23
2014-02-15
`
)

func TestParse(t *testing.T) {
	dates := Parse(testdata)
	result := ""

	for _, d := range dates {
		result = fmt.Sprintf("%s%s\n", result, d.Format("2006-01-02"))
	}

	if result != resultData {
		t.Errorf("got:\n`%s`\n\nexpected:\n`%s`\n\n", result, resultData)
	}
}
