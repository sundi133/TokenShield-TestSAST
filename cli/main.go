// Write config file
if err := viper.WriteConfig(); err != nil {
	// Config file doesn't exist, create it
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		os.Exit(1)
	}
	
	configPath := home + "/.tokenshield.yaml"
	viper.SetConfigFile(configPath)

	// --- FIX START ---
	f, ferr := os.OpenFile(configPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if ferr == nil {
		f.Close()
	}
	// --- FIX END ---

	if err := viper.WriteConfigAs(configPath); err != nil {
		fmt.Printf("Error creating config file: %v\n", err)
		fmt.Printf("Session saved temporarily but won't persist\n")
	} else {
		fmt.Printf("Created config file: %s\n", configPath)
		// Set secure permissions on new config file
		os.Chmod(configPath, 0600)
	}
} else {
	// Set secure permissions on config file after writing sensitive data
	configFile := viper.ConfigFileUsed()
	if configFile != "" {
		os.Chmod(configFile, 0600) // Fix: restrict config file permissions
	}
}
// 🔒 VOTAL.AI Security Fix: Potential Sensitive Data Exposure via Insecure Config File Permissions [CWE-312] - HIGH