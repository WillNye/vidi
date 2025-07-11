# Awesome Project

A real-time video processing pipeline.

## Features

- **Image Capture**
- **Flexible Pipeline**: Event-driven architecture for easy extension
- **Multiple Outputs**: Save both raw and processed frames
- **Concurrent Processing**: Handle multiple video streams efficiently

## Architecture

The project follows a modular, event-driven architecture with these main components:

### Core Components

- **Data Source**: Handles video input (currently webcam)
- **Data Classifier**: Processes frames (currently face detection)
- **Data Destination**: Handles output (file storage)

### Data Flow

```
Data Source → [Classifier] → Processed Frames → File Storage
     ↓
   Processed Frames → File Storage
```

## Installation

1. Install Go (1.16 or later)
2. Install OpenCV and its Go bindings:
   ```bash
   brew install opencv
   ```
3. Install project dependencies:
   ```bash
   go mod download
   ```

## Usage

1. Run the application:
   ```bash
   go run main.go
   ```

## Project Structure

```
.
├── core/                 # Core interfaces and types
├── data_classifier/      # Classification logic
├── data_destination/     # Output handlers (file storage)
└── data_source/          # Input handlers
```

## Dependencies

- [gocv.io/x/gocv](https://gocv.io/): Go bindings for OpenCV
- [github.com/asaskevich/EventBus](https://github.com/asaskevich/EventBus): Event bus for inter-component communication

