package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var dbsecret = "{\"password\":\"postgres\",\"dbname\":\"user_db\",\"port\":\"5432\",\"username\":\"postgres\",\"host\":\"localhost\"}"

func TestConfig_DBSecret(t *testing.T) {
	require := require.New(t)
	cfg := &Config{}
	cfg.DbSecrets = dbsecret

	err := cfg.parseDbSecrets()
	require.Nil(err)

	t.Log(cfg.PGConnectionString)
	require.Equal("port=5432 host=localhost user=postgres password=postgres dbname=user_db sslmode=disable", cfg.PGConnectionString)
}

func TestConfig_Validate(t *testing.T) {
	type fields struct {
		ListenAddr     string
		DefaultTimeout int
		Debug          bool
		Stage          string
		Branch         string
		DbSecrets      string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "all required values should validate",
			fields: fields{
				Stage:  "dev",
				Branch: "master",
			},
			wantErr: false,
		},
		{
			name: "missing stage should not validate",
			fields: fields{
				Branch: "master",
			},
			wantErr: true,
		},
		{
			name: "missing branch should not validate",
			fields: fields{
				Stage: "dev",
			},
			wantErr: true,
		},
		{
			name: "missing branch should not validate",
			fields: fields{
				Branch:    "master",
				Stage:     "dev",
				DbSecrets: dbsecret,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &Config{
				Port:   tt.fields.ListenAddr,
				Debug:  tt.fields.Debug,
				Stage:  tt.fields.Stage,
				Branch: tt.fields.Branch,
			}
			if err := cfg.validate(); (err != nil) != tt.wantErr {
				t.Errorf("Config.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
