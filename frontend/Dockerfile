FROM node:19
WORKDIR /app/frontend
COPY package*.json ./
RUN npm install
RUN npm install axios
RUN npm install react-router-dom
RUN npm install @mui/material @emotion/react @emotion/styled
COPY . .
CMD ["npm", "start"]