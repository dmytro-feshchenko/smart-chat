FROM node:7.8.0

# The base node image sets a very verbose log level.
ENV NPM_CONFIG_LOGLEVEL warn

# Create app directory
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

# Install app dependencies
COPY package.json /usr/src/app/
RUN npm install

# Bundle app source
COPY . /usr/src/app

RUN npm run build --production

# In your Dockerfile.
RUN npm install -g serve
# Run serve when the image is run.
CMD serve -s build
# Let Docker know about the port that serve runs on.
EXPOSE 5000
