# cobra-ui 

The cobra-ui package empowers developers to craft immersive and interactive user interfaces seamlessly into their CLI applications (cobra-ui is a cross platform UI that can be integrated easily with Cobra CLIs [Cobra](https://github.com/spf13/cobra) or other CLI tools )

## Features

- **File Input with Pagination**: Facilitate the selection of files from a directory with built-in pagination support (10 files/page). Users can navigate through large sets of files seamlessly.

- **Error Handling**: Handle errors gracefully during user interactions. When a handler encounters an error, it returns an error object containing relevant information about the error. The UI displays the error message to the user and prompts the question again.
  
- **Single-Choice Questions**: Allow users to select one option from a list of choices.
  
- **Text Input Questions**: Prompt users to enter text-based inputs.
  
- **Password Input Questions**: Securely collect password inputs from users, hiding the entered characters for privacy.
  
- **Dynamic Pagination**: Automatically paginate choices for single-choice questions with more than 10 options, ensuring a smooth user experience without overwhelming them with too many choices at once.
  
- Each question can have separately its own string cursor and color (the default cursor if not specified is ->)
