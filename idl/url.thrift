namespace go api

enum Code {
    Success         = 1,
    ParamInvalid    = 2,
    DBErr           = 3,
}

struct Url {
    1: i64 id
    2: string longUrl
    3: string shortUrl
    4: i64 userId
    5: i64 visits
    6: i32 rank
}

struct GenerateShortLinkRequest {
    1: string longUrl
}

struct GenerateShortLinkResponse {
    1: Code code
    2: string shortUrl
    3: string Message
}

struct DeleteShortLinkRequest {
    1: i64 id
}

struct DeleteShortLinkResponse {
    1: Code code
    2: string Message
}

struct UpdateShortLinkRequest {
    1: i64 id
    2: string oldShortUrl
}

struct UpdateShortLinkResponse {
    1: Code code
    2: string newShortUrl
    3: string Message
}

struct GetShortLinkDetailsRequest {
    1: i64 id
    2: i64 userId
}

struct GetUserShortLinksRequest {
    1: i64 userId
}

struct GetShortLinkRankingsRequest {}

service ShortLinkService {
    GenerateShortLinkResponse generateShortLink(1: GenerateShortLinkRequest req)
    DeleteShortLinkResponse deleteShortLink(1: DeleteShortLinkRequest req)
    UpdateShortLinkResponse updateShortLink(1: UpdateShortLinkRequest req)
    list<Url> getUserShortLinks(1: GetUserShortLinksRequest req)
    list<Url> getShortLinkRankings(1: GetShortLinkRankingsRequest req)
}
