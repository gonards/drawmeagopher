package main

// Server struct
type Server struct {
	Port string `yaml:"port"`
}

// Render struct
type Render struct {
	ImagePath string `yaml:"image_path"`
}

// Config struct
type Config struct {
	Server     Server     `yaml:"server"`
	Render     Render     `yaml:"render"`
	Categories []Category `yaml:"categories"`
}

// Image struct
type Image struct {
	Href          string `json:"href"`
	ID            string `json:"id"`
	Name          string `json:"name" yaml:"name"`
	ThumbnailHref string `json:"thumbnail_href"`
}

// Category struct
type Category struct {
	ID     string  `json:"id"`
	Images []Image `json:"images"`
	Name   string  `json:"name" yaml:"name"`
	Weight int     `yaml:"weight"`
}

// Artwork struct
type Artwork struct {
	TotalCombinations int        `json:"totalCombination"`
	Categories        []Category `json:"categories"`
}
