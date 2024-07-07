package parsing

// TerraformLock represents a .terraform.lock.hcl dependency lock file.
type TerraformLock struct {
	Providers []Provider `hcl:"provider,block"`
}

// Provider represents a provider block in a .terraform.lock.hcl file.
type Provider struct {
	Address     string   `hcl:"address,label"`
	Version     string   `hcl:"version"`
	Constraints string   `hcl:"constraints"`
	Hashes      []string `hcl:"hashes"`
}
