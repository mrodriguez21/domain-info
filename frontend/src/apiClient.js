const BASE_URI = "http://localhost:8090";

const APIClient = {
  getDomain(domainName) {
    return fetch(`${BASE_URI}/domains/${domainName}`).then((response) =>
      response.json()
    );
  },

  getPreviousDomains() {
    return fetch(`${BASE_URI}/domains`).then((response) => response.json());
  },
};

export default APIClient;
