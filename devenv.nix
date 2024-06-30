{ pkgs, ... }:
{
  languages.go.enable = true;

  packages = [ pkgs.coreutils pkgs.go-migrate pkgs.sqlc];

  services.postgres = {
    enable = true;
    initialDatabases = [{ name = "share_mail"; }];
  };

  pre-commit.hooks = {
    gofmt.enable = true;
    govet.enable = true;
    gotest.enable = true;
  };

  scripts = {
    db-migrate.exec = ''
      migrate -path database/migrations -database "postgresql:///share_mail?sslmode=disable" $@
    '';

    db-drop.exec = ''
      rm -r .devenv/state/postgres
    '';
  };

  env.CGO_ENABLED = 0;
}
