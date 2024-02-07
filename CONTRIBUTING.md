## Contribution Guide
Contributions from the community to enhance and improve this project are welcome. If you would like to contribute, please follow these guidelines:

1. **Fork the Repository:** Start by forking the repository to your own GitHub account.
2. **Clone the Repository:** Clone the forked repository to your local machine.
```bash
git clone https://github.com/yourusername/yourproject.git
```
3. **Create a Branch**: Create a new branch for your contributions.
```bash
git checkout -b <feature-or-fix-name>
```
4. **Make Changes**: Implement your fixture or fix and make sure the tests pass.
5. **Run Tests**: Ensure that your changes pass the existing tests. If you added new features, consider adding tests.
> There's also an opportunity to further improve and expand the existing test suite.
```bash
// Add your tests in this directory ./tests/. Please, ensure to follow existing patterns where possible.

go test ./tests/
```
6. **Commit and Push Changes**: Commit your changes with a descriptive commit message, push your changes to your forked repository.
```
git add .
git commit -m "Add feature or fix"
git push origin feature-or-fix-name
```
7. **Create a Pull Request**: Open a pull request on the original repository with a clear title and description of your changes.

If you'd prefer to just point out some missing features or suggest a fix to existing implementations, please go ahead and create a new issue with your suggestions.

Thank you!!!
