FROM nginx:alpine

# Purani default config delete karo
RUN rm /etc/nginx/conf.d/default.conf

# Apni secure config copy karo
COPY nginx.conf /etc/nginx/nginx.conf

# Nginx port 80 par chalta hai
EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]