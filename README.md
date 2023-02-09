# gopenai

[![codecov](https://codecov.io/gh/psyb0t/gopenai/branch/master/graph/badge.svg?token=NOYS3XR0BJ)](https://codecov.io/gh/psyb0t/gopenai)
[![goreportcard](https://goreportcard.com/badge/github.com/psyb0t/gopenai)](https://goreportcard.com/report/github.com/psyb0t/gopenai)
[![pipeline](https://github.com/psyb0t/gopenai/actions/workflows/pipeline.yml/badge.svg)](https://github.com/psyb0t/gopenai/actions/workflows/pipeline.yml)

gopenai is an unofficial package that provides bindings for the OpenAI API. It exposes a collection of interfaces for interacting with models, completions, edits, images, embeddings, files, fine tunes, and moderations.

## Config

To use the package, you'll need to create a `Config` struct with your OpenAI API key and organization ID. You can also set a custom request timeout.

```go
cfg := Config{
    APIKey: os.Getenv("OPENAI_API_KEY"),
    OrganizationID: "YOUR_ORGANIZATION_ID",
    RequestTimeout: time.Second * 30
}
```

## Client

You can create a new client that uses the config you created with `New()`.

```go
c := gopenai.New(cfg)
```

The client provides access to the different APIs.

```go
modelsAPI := c.Models()
completionsAPI := c.Completions()
editsAPI := c.Edits()
imagesAPI := c.Images()
embeddingsAPI := c.Embeddings()
filesAPI := c.Files()
fineTunesAPI := c.FineTunes()
moderationsAPI := c.Moderations()
```

## ModelsAPI

The Models API allows you to get a list of all models, get a model by ID, and delete a model by ID.

```go
models, err := modelsAPI.GetAll()
model, err := modelsAPI.GetByID(id)
deletedModel, err := modelsAPI.DeleteByID(id)
```

## CompletionsAPI

The Completions API allows you to create a completion.

```go
params := CompletionParams{
    Model: "YOUR_MODEL",
    Prompt: "YOUR_PROMPT",
    Suffix: "YOUR_SUFFIX",
    MaxTokens: 30,
    Temperature: 1.0,
    TopP: 0.9,
    N: 10,
    Logprobs: 1,
    Echo: true,
    Stop: ".",
    PresencePenalty: 0.0,
    FrequencyPenalty: 0.0,
    BestOf: 3,
    User: "YOUR_USER"
}

completion, err := completionsAPI.Create(params)
```

## EditsAPI

The Edits API provides methods for creating edits.

```go
params := EditParams{
    Model: "model_name",
    Input: "input_text",
    Instruction: "edit_instruction",
    N: 10,
    Temperature: 0.5,
    TopP: 0.8,
}

edit, err := editsAPI.Create(params)
```

# Images API

The Images API provides methods for creating, editing, and creating variations of images.

## Create

The `Create` method creates a new image based on the given input.

```go
params := gopenai.ImageGenerationParams{
    Prompt: "This is a test prompt",
    N: 10,
    Size: gopenai.ImageSize1024x1024,
    ResponseFormat: gopenai.ImageResponseFormatURL,
    User: "username",
}

images, err := imagesAPI.Create(params)
```

## Edit

The `Edit` method edits an existing image based on the given input.

```go
params := gopenai.ImageEditParams{
    Image: "path/to/image.png",
    Mask: "path/to/mask.png",
    Prompt: "This is a test prompt",
    N: 10,
    Size: gopenai.ImageSize1024x1024,
    ResponseFormat: gopenai.ImageResponseFormatURL,
    User: "username",
}

images, err := imagesAPI.Edit(params)
```

## Create Variations

The `CreateVariations` method creates variations of an existing image based on the given input.

```go
params := gopenai.ImageVariationParams{
    Image: "path/to/image.png",
    N: 10,
    Size: gopenai.ImageSize1024x1024,
    ResponseFormat: gopenai.ImageResponseFormatURL,
    User: "username",
}

images, err := imagesAPI.CreateVariations(params)
```

## Embeddings API

The Embeddings API provides methods to create embeddings from the OpenAI models.

### Create

The `Create` method takes an `EmbeddingParams` struct as a parameter and returns an `Embedding` struct.

```go
params := gopenai.EmbeddingParams{
    Model: "text-embedding-ada-002",
    Input: "This is a string to be embedded",
    User: "user@example.com",
}

embedding, err := embeddingsAPI.Create(params)
```

## Files API

The Files API provides methods to manage files, such as creating, deleting and downloading files.

### GetAll

The `GetAll` method returns a list of `File` structs.

```go
files, err := filesAPI.GetAll()
```

### GetByID

The `GetByID` method takes an ID as a parameter and returns a `File` struct.

```go
file, err := filesAPI.GetByID("<file_id>")
```

### Create

The `Create` method takes a `FileParams` struct as a parameter and returns a `File` struct.

```go
params := gopenai.FileParams{
    File: "<path_to_file>",
    Purpose: "fine-tune",
}

file, err := filesAPI.Create(params)
```

### DeleteByID

The `DeleteByID` method takes an ID as a parameter and returns a `DeletedFile` struct.

```go
deleted, err := filesAPI.DeleteByID("<file_id>")
```

### DownloadByID

The `DownloadByID` method takes an ID and a destination `io.Writer` as parameters and returns an error.

```go
err := filesAPI.DownloadByID("<file_id>", os.Stdout)
```

## FineTunes API

The FineTunes API provides methods to manage fine tunes, such as creating and cancelling fine tunes.

### GetAll

The `GetAll` method returns a list of `FineTune` structs.

```go
fineTunes, err := fineTunesAPI.GetAll()
```

### GetByID

The `GetByID` method takes an ID as a parameter and returns a `FineTune` struct.

```go
fineTune, err := fineTunesAPI.GetByID("<fine_tune_id>")
```

### Create

The `Create` method takes a `FineTuneParams` struct as a parameter and returns a `FineTune` struct.

```go
params := gopenai.FineTuneParams{
    TrainingFile: "<training_file_id>",
    ValidationFile: "<validation_file_id>",
    Model: "curie",
    NEpochs: 3,
    BatchSize: 16,
    LearningRateMultiplier: 0.1,
    PromptLossWeight: 0.5,
    ComputeClassificationMetrics: true,
    ClassificationNClasses: 2,
    ClassificationPositiveClass: "positive",
    ClassificationBetas: []float64{0.5, 0.5},
    Suffix: "my_suffix",
}

fineTune, err := fineTunesAPI.Create(params)
```

### Cancel

The `Cancel` method takes an ID as a parameter and returns a `FineTune` struct.

```go
fineTune, err := fineTunesAPI.Cancel("<fine_tune_id>")
```

### GetEvents

The `GetEvents` method takes a fine tune ID as a parameter and returns a list of `FineTuneEvent` structs.

```go
events, err := fineTunesAPI.GetEvents("<fine_tune_id>")
```

## Moderations API

The Moderations API provides methods to manage moderations, such as creating moderations.

### Create

The `Create` method takes a `ModerationParams` struct as a parameter and returns a `Moderation` struct.

```go
params := gopenai.ModerationParams{
    Input: "This is a string to be moderated",
    Model: "text-moderation-stable",
}

moderation, err := moderationsAPI.Create(params)
```

## TODO

- add more tests
