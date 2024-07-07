FROM node:19
WORKDIR /app/frontend
COPY package*.json ./
RUN npm install
RUN npm install react-router-dom
COPY . .
CMD ["npm", "start"]