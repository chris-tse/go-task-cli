# Dockerfile
FROM golang:1.23.3

# Avoid prompts during package installation
ENV DEBIAN_FRONTEND=noninteractive

# Install basic dependencies
RUN apt-get update && apt-get install -y \
    zsh \
    git \
    ripgrep \
    jq \
    wget \
    lsd \
    && rm -rf /var/lib/apt/lists/*

# Install starship
RUN curl -sS https://starship.rs/install.sh | sh -s -- --yes

# Install zellij
RUN wget https://github.com/zellij-org/zellij/releases/latest/download/zellij-x86_64-unknown-linux-musl.tar.gz \
    && tar xf zellij-x86_64-unknown-linux-musl.tar.gz \
    && mv zellij /usr/local/bin/ \
    && rm zellij-x86_64-unknown-linux-musl.tar.gz

#################################################################################
# Build and Install Neovim from Source
#################################################################################
ENV apt_update="apt-get update"
ENV apt_install="TERM=linux DEBIAN_FRONTEND=noninteractive apt-get install -q --yes --no-install-recommends"
ENV apt_clean="apt-get clean && apt-get autoremove -y && apt-get purge -y --auto-remove"
ENV curl="/usr/bin/curl --silent --show-error --tlsv1.2 --location"
ENV curl_github="/usr/bin/curl --silent --show-error --tlsv1.2 --request GET --url"
ENV dir_clean="\
  rm -rf \
  /var/lib/{apt,cache,log} \
  /usr/share/{doc,man,locale} \
  /var/cache/apt \
  /home/*/.cache \
  /root/.cache \
  /var/tmp/* \
  /tmp/* \
  "
# Build Packages
ARG BUILD_PKGS="\
make \
wget \
build-essential \
ninja-build \
gettext \
libtool \
libtool-bin \
autoconf \
automake \
cmake \
pkg-config \
unzip \
doxygen \
"

# Optional: replace vim with nvim
# && update-alternatives --install /usr/bin/vim vim /usr/local/bin/nvim 100 \
# && update-alternatives --set vim /usr/local/bin/nvim \
# && apt-get purge -y --auto-remove $(echo "${BUILD_PKGS}" | tr -d '\n' | sed 's/  */ /g') \
RUN echo \
  && export NAME="neovim" \
  && export TEST="nvim --version" \
  && export REPOSITORY="neovim/neovim" \
  && export VERSION="$(${curl_github} https://api.github.com/repos/${REPOSITORY}/releases/latest | jq --raw-output .tag_name)" \
  && echo "---------------------------------------------------------" \
  && echo "INFO[${NAME}] Building Neovim Version: ${VERSION}" \
  && echo "---------------------------------------------------------" \
  && ${apt_update} \
  && bash -c "${apt_install} ${BUILD_PKGS}" \
  && git clone --depth 1 --branch ${VERSION} https://github.com/${REPOSITORY}.git /tmp/neovim \
  && cd /tmp/neovim \
  && make CMAKE_BUILD_TYPE=Release \
  && make install \
  && rm -rf /tmp/neovim \
  && bash -c "${apt_clean}" \
  && ${dir_clean} \
  && ${TEST} \
  && echo


# Create non-root user
ARG USERNAME=developer
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && chsh -s $(which zsh) $USERNAME

USER $USERNAME
WORKDIR /home/$USERNAME

# Set zsh as default shell
SHELL ["/bin/zsh", "-c"]

# Keep container running
CMD ["tail", "-f", "/dev/null"]
