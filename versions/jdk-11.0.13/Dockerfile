FROM centos:7

WORKDIR /opt

RUN curl -L -o /usr/local/bin/dumb-init https://github.com/Yelp/dumb-init/releases/download/v1.2.2/dumb-init_1.2.2_amd64
RUN chmod +x /usr/local/bin/dumb-init

# Install AEM dependencies.
RUN yum install -y epel-release \
  && rpm -v --import http://li.nux.ro/download/nux/RPM-GPG-KEY-nux.ro \
  && rpm -Uvh http://li.nux.ro/download/nux/dextop/el7/x86_64/nux-dextop-release-0-5.el7.nux.noarch.rpm \
  && yum install -y \
    openssl.x86_64 openssl.i686 openssl-devel \
    cjkuni-ukai-fonts.noarch cjkuni-uming-fonts.noarch \
    nss-softokn-freebl.x86_64 nss-softokn-freebl.i686 \
    glibc.x86_64 glibc.i686 \
    libX11.x86_64 libX11.i686 \
    libxcb.x86_64 libxcb.i686 \
    libXau.x86_64 libXau.i686 \
    zlib.x86_64 zlib.i686 \
    libXext.x86_64 libXext.i686 \
    fontconfig.x86_64 fontconfig.i686 \
    expat.x86_64 expat.i686 \
    freetype.x86_64 freetype.i686 \
    freetype-devel.x86_64 freetype-devel.i686 \
    libSM.x86_64 libSM.i686 \
    libICE.x86_64 libICE.i686 \
    libuuid.x86_64 libuuid.i686 \
    libuuid.x86_64 libuuid.i686 \
    libXrandr.x86_64 libXrandr.i686 \
    libXrender.x86_64 libXrender.i686 \
    libXinerama.x86_64 libXinerama.i686 \
    fontconfig.x86_64 fontconfig.i686 \
    ffmpeg \
    ImageMagick \
    ghostscript \
    zip \
    unzip \
  && yum clean all \
  && rm -rf /var/cache/yum

COPY .artifacts/jdk/jdk-11.0.13.rpm jdk.rpm
RUN yum install -y jdk.rpm
RUN rm jdk.rpm

# A volume for the AEM crx-quickstart directory and license file. Allows AEM data to persist across different containers.
VOLUME /opt/aem/

# AEM will run under the aem user and group.
RUN useradd -m -s /bin/bash -u 1000 aem

# Add a label for filtering in scripts/shell.
LABEL name="cord-base"

# Set the AEM environment variables.
ENV CQ_JVM_OPTS="-server -Xmx1024m -XX:MaxPermSize=256M -Djava.awt.headless=true -agentlib:jdwp=transport=dt_socket,server=y,suspend=n,address=*:1044 -XX:+UseParallelGC --add-opens=java.desktop/com.sun.imageio.plugins.jpeg=ALL-UNNAMED --add-opens=java.base/sun.net.www.protocol.jrt=ALL-UNNAMED --add-opens=java.naming/javax.naming.spi=ALL-UNNAMED --add-opens=java.xml/com.sun.org.apache.xerces.internal.dom=ALL-UNNAMED --add-opens=java.base/java.lang=ALL-UNNAMED --add-opens=java.base/jdk.internal.loader=ALL-UNNAMED --add-opens=java.base/java.net=ALL-UNNAMED -Dnashorn.args=--no-deprecation-warning"

# Expose the standard AEM ports e.g. 4502, 4503.

EXPOSE 4502

EXPOSE 4503

COPY scripts/start-aem.sh start-aem.sh
COPY scripts/stop_aem.sh stop_aem.sh

# Make the aem user the owner of the aem and backups volumes. 
RUN chown -R aem:aem /opt/

USER aem:aem

ENTRYPOINT ["/usr/local/bin/dumb-init", "--"]
CMD ["bash", "-c", "./start-aem.sh"]

