package country

import "strings"

/*
countrydata is derived from the ISO 3166-1 information
presented on wikipedia:
http://en.wikipedia.org/wiki/ISO_3166-1
The data was obtained from the wiki source for "Officially
assigned code elements". Some munging occurred, then the
tab-delimited csv file data in this source file was constructed.
*/
var countrydata = strings.NewReader(`Afghanistan	AF	AFG	004
Åland Islands	AX	ALA	248
Albania	AL	ALB	008
Algeria	DZ	DZA	012
American Samoa	AS	ASM	016
Andorra	AD	AND	020
Angola	AO	AGO	024
Anguilla	AI	AIA	660
Antarctica	AQ	ATA	010
Antigua and Barbuda	AG	ATG	028
Argentina	AR	ARG	032
Armenia	AM	ARM	051
Aruba	AW	ABW	533
Australia	AU	AUS	036
Austria	AT	AUT	040
Azerbaijan	AZ	AZE	031
Bahamas	BS	BHS	044
Bahrain	BH	BHR	048
Bangladesh	BD	BGD	050
Barbados	BB	BRB	052
Belarus	BY	BLR	112
Belgium	BE	BEL	056
Belize	BZ	BLZ	084
Benin	BJ	BEN	204
Bermuda	BM	BMU	060
Bhutan	BT	BTN	064
Bolivia, Plurinational State of	BO	BOL	068
Bonaire, Sint Eustatius and Saba	BQ	BES	535
Bosnia and Herzegovina	BA	BIH	070
Botswana	BW	BWA	072
Bouvet Island	BV	BVT	074
Brazil	BR	BRA	076
British Indian Ocean Territory	IO	IOT	086
Brunei Darussalam	BN	BRN	096
Bulgaria	BG	BGR	100
Burkina Faso	BF	BFA	854
Burundi	BI	BDI	108
Cambodia	KH	KHM	116
Cameroon	CM	CMR	120
Canada	CA	CAN	124
Cabo Verde	CV	CPV	132
Cayman Islands	KY	CYM	136
Central African Republic	CF	CAF	140
Chad	TD	TCD	148
Chile	CL	CHL	152
China	CN	CHN	156
Christmas Island	CX	CXR	162
Cocos (Keeling) Islands	CC	CCK	166
Colombia	CO	COL	170
Comoros	KM	COM	174
Republic of the Congo	CG	COG	178
Congo, the Democratic Republic of the	CD	COD	180
Cook Islands	CK	COK	184
Costa Rica	CR	CRI	188
Côte d'Ivoire	CI	CIV	384
Croatia	HR	HRV	191
Cuba	CU	CUB	192
Curaçao	CW	CUW	531
Cyprus	CY	CYP	196
Czech Republic	CZ	CZE	203
Denmark	DK	DNK	208
Djibouti	DJ	DJI	262
Dominica	DM	DMA	212
Dominican Republic	DO	DOM	214
Ecuador	EC	ECU	218
Egypt	EG	EGY	818
El Salvador	SV	SLV	222
Equatorial Guinea	GQ	GNQ	226
Eritrea	ER	ERI	232
Estonia	EE	EST	233
Ethiopia	ET	ETH	231
Falkland Islands (Malvinas)	FK	FLK	238
Faroe Islands	FO	FRO	234
Fiji	FJ	FJI	242
Finland	FI	FIN	246
France	FR	FRA	250
French Guiana	GF	GUF	254
French Polynesia	PF	PYF	258
French Southern Territories	TF	ATF	260
Gabon	GA	GAB	266
Gambia	GM	GMB	270
Georgia (country)|Georgia	GE	GEO	268
Germany	DE	DEU	276
Ghana	GH	GHA	288
Gibraltar	GI	GIB	292
Greece	GR	GRC	300
Greenland	GL	GRL	304
Grenada	GD	GRD	308
Guadeloupe	GP	GLP	312
Guam	GU	GUM	316
Guatemala	GT	GTM	320
Guernsey	GG	GGY	831
Guinea	GN	GIN	324
Guinea-Bissau	GW	GNB	624
Guyana	GY	GUY	328
Haiti	HT	HTI	332
Heard Island and McDonald Islands	HM	HMD	334
Vatican City State|Holy See (Vatican City State)	VA	VAT	336
Honduras	HN	HND	340
Hong Kong	HK	HKG	344
Hungary	HU	HUN	348
Iceland	IS	ISL	352
India	IN	IND	356
Indonesia	ID	IDN	360
Iran, Islamic Republic of	IR	IRN	364
Iraq	IQ	IRQ	368
Republic of Ireland|Ireland	IE	IRL	372
Isle of Man	IM	IMN	833
Israel	IL	ISR	376
Italy	IT	ITA	380
Jamaica	JM	JAM	388
Japan	JP	JPN	392
Jersey	JE	JEY	832
Jordan	JO	JOR	400
Kazakhstan	KZ	KAZ	398
Kenya	KE	KEN	404
Kiribati	KI	KIR	296
Korea, Democratic People's Republic of	KP	PRK	408
Korea, Republic of	KR	KOR	410
Kuwait	KW	KWT	414
Kyrgyzstan	KG	KGZ	417
Lao People's Democratic Republic	LA	LAO	418
Latvia	LV	LVA	428
Lebanon	LB	LBN	422
Lesotho	LS	LSO	426
Liberia	LR	LBR	430
Libya	LY	LBY	434
Liechtenstein	LI	LIE	438
Lithuania	LT	LTU	440
Luxembourg	LU	LUX	442
Macao	MO	MAC	446
Republic of Macedonia|Macedonia, the former Yugoslav Republic of	MK	MKD	807
Madagascar	MG	MDG	450
Malawi	MW	MWI	454
Malaysia	MY	MYS	458
Maldives	MV	MDV	462
Mali	ML	MLI	466
Malta	MT	MLT	470
Marshall Islands	MH	MHL	584
Martinique	MQ	MTQ	474
Mauritania	MR	MRT	478
Mauritius	MU	MUS	480
Mayotte	YT	MYT	175
Mexico	MX	MEX	484
Micronesia, Federated States of	FM	FSM	583
Moldova, Republic of	MD	MDA	498
Monaco	MC	MCO	492
Mongolia	MN	MNG	496
Montenegro	ME	MNE	499
Montserrat	MS	MSR	500
Morocco	MA	MAR	504
Mozambique	MZ	MOZ	508
Myanmar	MM	MMR	104
Namibia	NA	NAM	516
Nauru	NR	NRU	520
Nepal	NP	NPL	524
Netherlands	NL	NLD	528
New Caledonia	NC	NCL	540
New Zealand	NZ	NZL	554
Nicaragua	NI	NIC	558
Niger	NE	NER	562
Nigeria	NG	NGA	566
Niue	NU	NIU	570
Norfolk Island	NF	NFK	574
Northern Mariana Islands	MP	MNP	580
Norway	NO	NOR	578
Oman	OM	OMN	512
Pakistan	PK	PAK	586
Palau	PW	PLW	585
State of Palestine|Palestine, State of	PS	PSE	275
Panama	PA	PAN	591
Papua New Guinea	PG	PNG	598
Paraguay	PY	PRY	600
Peru	PE	PER	604
Philippines	PH	PHL	608
Pitcairn	PN	PCN	612
Poland	PL	POL	616
Portugal	PT	PRT	620
Puerto Rico	PR	PRI	630
Qatar	QA	QAT	634
Réunion	RE	REU	638
Romania	RO	ROU	642
Russian Federation	RU	RUS	643
Rwanda	RW	RWA	646
Saint Barthélemy	BL	BLM	652
Saint Helena, Ascension and Tristan da Cunha	SH	SHN	654
Saint Kitts and Nevis	KN	KNA	659
Saint Lucia	LC	LCA	662
Saint Martin (French part)	MF	MAF	663
Saint Pierre and Miquelon	PM	SPM	666
Saint Vincent and the Grenadines	VC	VCT	670
Samoa	WS	WSM	882
San Marino	SM	SMR	674
Sao Tome and Principe	ST	STP	678
Saudi Arabia	SA	SAU	682
Senegal	SN	SEN	686
Serbia	RS	SRB	688
Seychelles	SC	SYC	690
Sierra Leone	SL	SLE	694
Singapore	SG	SGP	702
Sint Maarten (Dutch part)	SX	SXM	534
Slovakia	SK	SVK	703
Slovenia	SI	SVN	705
Solomon Islands	SB	SLB	090
Somalia	SO	SOM	706
South Africa	ZA	ZAF	710
South Georgia and the South Sandwich Islands	GS	SGS	239
South Sudan	SS	SSD	728
Spain	ES	ESP	724
Sri Lanka	LK	LKA	144
Sudan	SD	SDN	729
Suriname	SR	SUR	740
Svalbard and Jan Mayen	SJ	SJM	744
Swaziland	SZ	SWZ	748
Sweden	SE	SWE	752
Switzerland	CH	CHE	756
Syrian Arab Republic	SY	SYR	760
Taiwan	TW	TWN	158
Tajikistan	TJ	TJK	762
Tanzania, United Republic of	TZ	TZA	834
Thailand	TH	THA	764
Timor-Leste	TL	TLS	626
Togo	TG	TGO	768
Tokelau	TK	TKL	772
Tonga	TO	TON	776
Trinidad and Tobago	TT	TTO	780
Tunisia	TN	TUN	788
Turkey	TR	TUR	792
Turkmenistan	TM	TKM	795
Turks and Caicos Islands	TC	TCA	796
Tuvalu	TV	TUV	798
Uganda	UG	UGA	800
Ukraine	UA	UKR	804
United Arab Emirates	AE	ARE	784
United Kingdom	GB	GBR	826
United States	US	USA	840
United States Minor Outlying Islands	UM	UMI	581
Uruguay	UY	URY	858
Uzbekistan	UZ	UZB	860
Vanuatu	VU	VUT	548
Venezuela, Bolivarian Republic of	VE	VEN	862
Viet Nam	VN	VNM	704
Virgin Islands, British	VG	VGB	092
Virgin Islands, U.S.	VI	VIR	850
Wallis and Futuna	WF	WLF	876
Western Sahara	EH	ESH	732
Yemen	YE	YEM	887
Zambia	ZM	ZMB	894
Zimbabwe	ZW	ZWE	716`)
