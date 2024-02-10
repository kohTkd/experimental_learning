ARG RUBY_VERSION=3.2.1
FROM ruby:${RUBY_VERSION}-alpine

# USER root

ENV APP_ROOT /app/

ENV BUILD_DEPS 'build-base alpine-sdk'
ENV EXEC_DEPS 'bash tzdata postgresql-dev libxslt-dev libxml2-dev'
ENV IMAGE_DEPS 'imagemagick imagemagick-dev'
ENV DEV_DEPS 'less'

RUN mkdir -p $APP_ROOT
# RUN adduser -h $APP_ROOT -D rails && \
#     mkdir -p $APP_ROOT && chown -R rails:rails $APP_ROOT

WORKDIR $APP_ROOT

RUN apk update && apk upgrade && \
    apk add --no-cache --virtual build-dependencies --update ${BUILD_DEPS} ${EXEC_DEPS} ${DEV_DEPS} && \
    rm -rf /var/cache/apk/*

# USER rails

COPY Gemfile Gemfile.lock $APP_ROOT

RUN gem install bundler && \
    bundle config set path 'vendor/bundle' && \
    bundle config set force_ruby_platform true && \
    bundle config --global build.nokogiri --use-system-libraries && \
    bundle install -j4

COPY . $APP_ROOT

EXPOSE 3000
EXPOSE 7658

CMD ["bundle", "exec", "rails", "s", "-p", "3000", "-b", "0.0.0.0"]
