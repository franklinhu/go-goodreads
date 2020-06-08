package goodreads

type Response struct {
	User    User     `xml:"user"`
	Book    Book     `xml:"book"`
	Reviews []Review `xml:"reviews>review"`
}

type AuthorResponse struct {
	Author Author `xml:"author"`
}

type SearchResponse struct {
	Results []Work `xml:"search>results>work"`
}

type User struct {
	ID            string       `xml:"id"`
	Name          string       `xml:"name"`
	About         string       `xml:"about"`
	Link          string       `xml:"link"`
	ImageURL      string       `xml:"image_url"`
	SmallImageURL string       `xml:"small_image_url"`
	Location      string       `xml:"location"`
	LastActive    string       `xml:"last_active"`
	ReviewCount   int          `xml:"reviews_count"`
	Statuses      []UserStatus `xml:"user_statuses>user_status"`
	Shelves       []Shelf      `xml:"user_shelves>user_shelf"`
	LastRead      []Review
}

func (u User) ReadingShelf() Shelf {
	for _, shelf := range u.Shelves {
		if shelf.Name == "currently-reading" {
			return shelf
		}
	}

	return Shelf{}
}

func (u User) ReadShelf() Shelf {
	for _, shelf := range u.Shelves {
		if shelf.Name == "read" {
			return shelf
		}
	}

	return Shelf{}
}

func (u User) ToReadShelf() Shelf {
	for _, shelf := range u.Shelves {
		if shelf.Name == "to-read" {
			return shelf
		}
	}

	return Shelf{}
}

type Shelf struct {
	ID        string `xml:"id"`
	BookCount string `xml:"book_count"`
	Name      string `xml:"name"`
}

type UserStatus struct {
	Page    int    `xml:"page"`
	Percent int    `xml:"percent"`
	Updated string `xml:"updated_at"`
	Book    Book   `xml:"book"`
}

func (u UserStatus) UpdatedRelative() string {
	return relativeDate(u.Updated)
}

type Book struct {
	ID       string   `xml:"id"`
	Title    string   `xml:"title"`
	Link     string   `xml:"link"`
	ImageURL string   `xml:"image_url"`
	NumPages string   `xml:"num_pages"`
	Format   string   `xml:"format"`
	Authors  []Author `xml:"authors>author"`
	ISBN     string   `xml:"isbn"`
}

func (b Book) Author() Author {
	return b.Authors[0]
}

type Author struct {
	ID                   string `xml:"id"`
	Name                 string `xml:"name"`
	Link                 string `xml:"link"`
	FansCount            int    `xml:"fans_count"`
	AuthorFollowersCount int    `xml:"author_followers_count"`
	LargeImageURL        string `xml:"large_image_url"`
	ImageURL             string `xml:"image_url"`
	SmallImageURL        string `xml:"small_image_url"`
	WorksCount           int    `xml:"works_count"`
	Gender               string `xml:"gender"`
	Hometown             string `xml:"hometown"`
}

type Review struct {
	Book   Book   `xml:"book"`
	Rating int    `xml:"rating"`
	ReadAt string `xml:"read_at"`
	Link   string `xml:"link"`
}

func (r Review) FullStars() []bool {
	return make([]bool, r.Rating)
}

func (r Review) EmptyStars() []bool {
	return make([]bool, 5-r.Rating)
}

func (r Review) ReadAtShort() string {
	date, err := parseDate(r.ReadAt)
	if err != nil {
		return ""
	}

	return (string)(date.Format("2 Jan 2006"))
}

func (r Review) ReadAtRelative() string {
	return relativeDate(r.ReadAt)
}

// `Work` is returned from search results and is an aggregate
// over many books
type Work struct {
	ID                       string  `xml:"id"`
	BooksCount               int     `xml:"books_count"`
	RatingsCount             int     `xml:"ratings_count"`
	TextReviewsCount         int     `xml:"text_reviews_count"`
	OriginalPublicationYear  int     `xml:"original_publication_year"`
	OriginalPublicationMonth int     `xml:"original_publication_month"`
	OriginalPublicationDay   int     `xml:"original_publication_day"`
	AverageRating            float32 `xml:"average_rating"`
	BestBook                 Book    `xml:"best_book"`
}
