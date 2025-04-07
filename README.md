
# Anime Recommendations (Go)

This project uses the `Anime Dataset 2023` to get information about different anime and make recommendations using the different data points present in the dataset.


## Run Locally

If you haven't already, download the dataset files from [here](https://www.kaggle.com/datasets/dbdmobile/myanimelist-dataset)

Clone the project

```bash
  git clone https://github.com/JalenMurray/AnimeRecommendationsGo
```

Go to the project directory

```bash
  cd AnimeRecommendationsGo
```

Build Executable
```bash
  go build
```

If you haven't already

Load datasets into db (This will take around 30 minutes)
```bash
  ./AnimeRecommendationsGo load-data
```

*Note:* If you don't have enough RAM to run this, you can lower `workerCount` in `db/loader.go` on Line 81.\
This will cause it to take longer to run though

Once all the data is loaded, you can start the REST server

Start REST server
```bash
  ./AnimeRecommendationsGo
```



## API Reference

#### Get anime

```http
  GET /anime/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of anime to fetch |

#### Get list of anime by query

```http
  GET /anime
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `name` | `string` | **Optional** Name to search for |
| `genres` | `string` | **Optional** Genres to search for |
| `score` | `number` | **Optional** Exact score match |
| `score_gt` | `number` | **Optional** Greater than score match |
| `score_lt` | `number` | **Optional** Less than score match |
| `episodes` | `number` | **Optional** Exact episodes match |
| `episodes_gt` | `number` | **Optional** Greater than episodes match |
| `episodes_lt` | `number` | **Optional** Less than episodes match |
| `popularity` | `number` | **Optional** Exact popularity match |
| `popularity_gt` | `number` | **Optional** Greater than popularity match |
| `popularity_lt` | `number` | **Optional** Less than popularity match |




## Acknowledgements

 - [Kaggle Dataset](https://www.kaggle.com/datasets/dbdmobile/myanimelist-dataset)


## License

[MIT](https://choosealicense.com/licenses/mit/)

