package models

type Artist struct {
    ID           int
    Name         string
    Image        string
    Members      []string
    CreationDate int
    FirstAlbum   string
    Locations    string 
    ConcertDates string 
    Relations    string 
}

type LocationEntry struct {
    ID        int
    Locations []string
    Dates     string
}

type DateEntry struct {
    ID    int
    Dates []string
}

type RelationEntry struct {
    ID             int
    DatesLocations map[string][]string
}

type FullArtist struct {
    ID             int
    Name           string
    Image          string
    Members        []string
    CreationDate   int
    FirstAlbum     string
    Locations      []string
    ConcertDates   []string
    DatesLocations map[string][]string
}

type LocationsResponse struct {
    Index []LocationEntry
}