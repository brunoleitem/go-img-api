# Image Watermark API

This project was created as a learning exercise in **Go**, inspired by a real feature I implemented in a production environment using **C#**.

This API allows users to upload an image, add a custom text watermark, and retrieve the watermarked image and a passkey to delete if needed.
It is built using **Go**, integrates with **Cloudflare R2** for storage, and uses **Redis** for memory.

## Features

- **Image Upload**: Accepts image uploads and validates the file type. Currently supported: PNG, JPG, JPEG.
- **Watermarking**: Adds a custom text watermark to the image.
- **Cloud Storage**: Stores watermarked images in **Cloudflare R2**.
- **Caching**: Utilizes **Redis** to store a user passkey to delete image.

## Endpoints

- `POST /upload`: Uploads an image and applies the watermark. Passing a query "text" with the user input text.
- `DELETE /image/{userPassKey}`: Deletes an image from storage and cache.

## Technologies

- **Go**
- **Cloudflare R2**
- **Redis**

## Setup Instructions

1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/image-watermark-api.git
   cd image-watermark-api
   ```
   
2. Install dependencies:
    ```bash
    go mod download
    ```

3. Set up environment variables (can be found in env.example)
4. Run server:
    ```bash
    go run cmd/server/main.go
    ```

## ToDo
[ ] Cache images with redis
[ ] Front end view
