# Crossplane Manifest Generator

---

### Features
- Create Crossplane provider manifests from a list read from a configs.yaml file
- more to come! Stay tuned

---

### Usage:
- clone this repo
- go mod init/tidy as needed
- Modify the templates/configs.yaml file for the providers you want to create (
  see https://marketplace.upbound.io/providers for a full list )
- Once you update the file, go run the main.go file.
- The program will read the configs, generate provider files, and put them in the "providers" folder.
- you can then install them into your crossplane management kubernetes cluster.