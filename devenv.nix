{ pkgs, ... }:
{
  languages.go.enable = true;

  packages = [ pkgs.coreutils pkgs.go-migrate ];
  services.postgres = {
    enable = true;
    initialDatabases = [{ name = "share_mail"; }];
  };

  scripts = {
    db-migrate.exec = ''
      migrate -path database/migrations -database "postgresql:///share_mail?sslmode=disable" $@
    '';

    db-drop.exec = ''
      devenv processes stop && rm -r .devenv/state/postgres && devenv up -d
    '';
  };

  env.CGO_ENABLED = 0;
}
