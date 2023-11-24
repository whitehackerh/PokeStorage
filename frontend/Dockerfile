FROM node:19
WORKDIR /app/frontend
COPY package*.json ./
RUN npm install
COPY . .
CMD ["npm", "start"]