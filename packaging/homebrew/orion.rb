class O < Formula
  desc "Natural-language terminal assistant"
  homepage "https://github.com/YOUR_USERNAME/orion"
  version "0.1.0"

  on_macos do
    if Hardware::CPU.intel?
      # URL to your amd64 binary from GitHub Releases
      url "https://github.com/YOUR_USERNAME/orion/releases/download/v0.1.0/orion_darwin_amd64"
      # Run 'shasum -a 256 dist/orion_darwin_amd64' to get this
      sha256 "REPLACE_WITH_AMD64_SHA256"
    elsif Hardware::CPU.arm?
      # URL to your arm64 binary from GitHub Releases
      url "https://github.com/YOUR_USERNAME/orion/releases/download/v0.1.0/orion_darwin_arm64"
      # Run 'shasum -a 256 dist/orion_darwin_arm64' to get this
      sha256 "REPLACE_WITH_ARM64_SHA256"
    end
  end

  def install
    # Rename binary to 'o' and install to bin
    if Hardware::CPU.intel?
      bin.install "orion_darwin_amd64" => "o"
    else
      bin.install "orion_darwin_arm64" => "o"
    end
  end

  test do
    system "#{bin}/o", "--help"
  end
end
