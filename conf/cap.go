package conf

var AccessoryCapList = []Accessory{
	{
		Id:     1,
		Name:   "Army Hat",
		Chance: 0.02250,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Army Hat",
			},
		},
	},
	{
		Id:     2,
		Name:   "Baby's Bonnet",
		Chance: 0.0240,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Baby's Bonnet",
			},
		},
	},
	{
		Id:     3,
		Name:   "Beanie",
		Chance: 0.02380,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Beanie",
			},
		},
	},
	{
		Id:     4,
		Name:   "Black Baseball Cap",
		Chance: 0.02360,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Black Baseball Cap",
			},
		},
	},
	{
		Id:     5,
		Name:   "Black Cap",
		Chance: 0.0300,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Black Cap",
			},
		},
	},
	{
		Id:     6,
		Name:   "Brown cap",
		Chance: 0.02320,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Brown cap",
			},
		},
	},
	{
		Id:     7,
		Name:   "Cowboy Hat",
		Chance: 0.0230,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Cowboy Hat",
			},
		},
	},
	{
		Id:     8,
		Name:   "Fashion Hat",
		Chance: 0.02250,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Fashion Hat",
			},
		},
	},
	{
		Id:     9,
		Name:   "Fedora Hat",
		Chance: 0.0220,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Fedora Hat",
			},
		},
	},
	{
		Id:     10,
		Name:   "Fez",
		Chance: 0.02150,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Fez",
			},
		},
	},
	{
		Id:     11,
		Name:   "Fisherman's Hat",
		Chance: 0.021,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Fisherman's Hat",
			},
		},
	},
	{
		Id:     12,
		Name:   "Party Hat",
		Chance: 0.02,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Party Hat",
			},
		},
	},
	{
		Id:     13,
		Name:   "Safari",
		Chance: 0.02,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Safari",
			},
		},
	},
	{
		Id:     14,
		Name:   "Seaman's Hat",
		Chance: 0.019,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Seaman's Hat",
			},
		},
	},
	{
		Id:     15,
		Name:   "Service Hat",
		Chance: 0.018,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Service Hat",
			},
		},
	},
	{
		Id:     16,
		Name:   "Aviator's helmet",
		Chance: 0.015,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Aviator's helmet",
			},
		},
	},
	{
		Id:     17,
		Name:   "Blue Turban",
		Chance: 0.014,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Blue Turban",
			},
		},
	},
	{
		Id:     18,
		Name:   "Laurel Wreath",
		Chance: 0.013,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Laurel Wreath",
			},
		},
	},
	{
		Id:     19,
		Name:   "Pigtail",
		Chance: 0.012,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Pigtail",
			},
		},
	},
	{
		Id:     20,
		Name:   "Red Baseball Cap",
		Chance: 0.011,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Red Baseball Cap",
			},
		},
	},
	{
		Id:     21,
		Name:   "S&M Hat",
		Chance: 0.01,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "S&M Hat",
			},
		},
	},
	{
		Id:     22,
		Name:   "Winning Headband",
		Chance: 0.009,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Winning Headband",
			},
		},
	},
	{
		Id:     23,
		Name:   "Witch Hat",
		Chance: 0.00800,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Witch Hat",
			},
		},
		IncompatibleAccessories: map[AccessoryTypeName][]int{
			AccessoryTypeNameHorn: []int{8},
		},
	},
	{
		Id:     24,
		Name:   "Blue Hair",
		Chance: 0.0038,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Blue Hair",
			},
		},
	},
	{
		Id:     25,
		Name:   "Gold Hair",
		Chance: 0.0038,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Gold Hair",
			},
		},
	},
	{
		Id:     26,
		Name:   "Green Hair",
		Chance: 0.0038,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Green Hair",
			},
		},
	},
	{
		Id:     27,
		Name:   "Prussian Helmet",
		Chance: 0.0038,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Prussian Helmet",
			},
		},
	},
	{
		Id:     28,
		Name:   "Purple Hair",
		Chance: 0.0038,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Purple Hair",
			},
		},
	},
	{
		Id:     29,
		Name:   "Red Hair",
		Chance: 0.0038,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Red Hair",
			},
		},
	},
	{
		Id:     30,
		Name:   "Sea Captain's Hat",
		Chance: 0.0038,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Sea Captain's Hat",
			},
		},
	},
	{
		Id:     31,
		Name:   "Stuntman Helmet",
		Chance: 0.0038,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Stuntman Helmet",
			},
		},
	},
	{
		Id:     32,
		Name:   "Color cap",
		Chance: 0.00180,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Color cap",
			},
		},
	},
	{
		Id:     33,
		Name:   "Colorful Hair",
		Chance: 0.00170,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Colorful Hair",
			},
		},
	},
	{
		Id:     34,
		Name:   "Crown",
		Chance: 0.00160,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Crown",
			},
		},
	},
	{
		Id:     35,
		Name:   "Spinner Hat",
		Chance: 0.00150,
		PropertyList: []Property{
			{
				Type: PropertyTypeNameCap,
				Name: "Spinner Hat",
			},
		},
	},
	//{
	//	Id:     36,
	//	Name:   "astronaut helmet",
	//	Chance: 0.000625,
	//	PropertyList: []Property{
	//		{
	//			Type: PropertyTypeNameCap,
	//			Name: "astronaut helmet",
	//		},
	//	},
	//},
}
