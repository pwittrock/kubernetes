# Builds slate md source into static html files
# For example source files see the "source" directory in the slate github repo
# docker run -ti -v /absolute/path/to/source:/slate/source -v /absolute/path/to/build:/slate/build pwittrock/slate

FROM ruby
MAINTAINER Phillip Wittrock <pwittroc@google.com>

RUN git clone --depth=1 https://github.com/lord/slate.git
WORKDIR slate

RUN gem install bundle
RUN bundle install

RUN apt-get update && apt-get install -y nodejs && apt-get clean && rm -rf /var/lib/apt/lists/*

CMD ["bundle", "exec", "middleman", "build", "--clean"]