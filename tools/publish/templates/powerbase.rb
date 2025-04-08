# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class {{ .FormulaName }} < Formula
  desc "{{ .Description }}"
  homepage "https://powerbase.club"
  version "{{ .Version }}"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/skorpland/cli/releases/download/v{{ .Version }}/powerbase_darwin_arm64.tar.gz"
      sha256 "{{ .Checksum.powerbase_darwin_arm64 }}"

      def install
        bin.install "powerbase"
        (bash_completion/"powerbase").write `#{bin}/powerbase completion bash`
        (fish_completion/"powerbase.fish").write `#{bin}/powerbase completion fish`
        (zsh_completion/"_powerbase").write `#{bin}/powerbase completion zsh`
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/skorpland/cli/releases/download/v{{ .Version }}/powerbase_darwin_amd64.tar.gz"
      sha256 "{{ .Checksum.powerbase_darwin_amd64 }}"

      def install
        bin.install "powerbase"
        (bash_completion/"powerbase").write `#{bin}/powerbase completion bash`
        (fish_completion/"powerbase.fish").write `#{bin}/powerbase completion fish`
        (zsh_completion/"_powerbase").write `#{bin}/powerbase completion zsh`
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/skorpland/cli/releases/download/v{{ .Version }}/powerbase_linux_arm64.tar.gz"
      sha256 "{{ .Checksum.powerbase_linux_arm64 }}"

      def install
        bin.install "powerbase"
        (bash_completion/"powerbase").write `#{bin}/powerbase completion bash`
        (fish_completion/"powerbase.fish").write `#{bin}/powerbase completion fish`
        (zsh_completion/"_powerbase").write `#{bin}/powerbase completion zsh`
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/skorpland/cli/releases/download/v{{ .Version }}/powerbase_linux_amd64.tar.gz"
      sha256 "{{ .Checksum.powerbase_linux_amd64 }}"

      def install
        bin.install "powerbase"
        (bash_completion/"powerbase").write `#{bin}/powerbase completion bash`
        (fish_completion/"powerbase.fish").write `#{bin}/powerbase completion fish`
        (zsh_completion/"_powerbase").write `#{bin}/powerbase completion zsh`
      end
    end
  end
end
