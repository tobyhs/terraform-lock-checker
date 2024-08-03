# terraform-lock-checker

terraform-lock-checker is a program that checks that .terraform.lock.hcl files have the expected number of provider hashes/platforms.

For example, you may want the .terraform.lock.hcl files in a repo to be created with `terraform providers lock -platform=linux_amd64 -platform=darwin_amd64 -platform=darwin_arm64 -platform=windows_amd64`. To check that the .terraform.lock.hcl files have 4 hashes for each provider, you can run `terraform-lock-checker -expected-count 4`. 

## Example Usage

```
$ terraform-lock-checker -expected-count 4 -exclusion-patterns 'test/data/invalid/*'
The following files have an unexpected number of provider checksums:
test/data/hash_1/.terraform.lock.hcl
```

## Container Image

A container image is pushed to ghcr.io when a release is published.

Example usage of the image:
```
$ docker run --rm -v "${PWD}:/workspace" ghcr.io/tobyhs/terraform-lock-checker:latest -exclusion-patterns 'test/data/invalid/*'
The following files have an unexpected number of provider checksums:
test/data/.terraform.lock.hcl
test/data/hash_4/.terraform.lock.hcl
```
