import url from 'url';

const getUrl = (relativeUrl) => {
  switch (process.env.NODE_ENV) {
    case "production": return url.resolve("http://127.0.0.1:9000/", relativeUrl);
    case "development": return relativeUrl;
    default: return relativeUrl;
  }
}
export const get = async (url) => {
  const response = await fetch(getUrl(url));
  return await parseResponse(response);
};

export const post = async (url, body) => {
  const response = await fetch(getUrl(url), { method: "POST", body: JSON.stringify(body) });
  return await parseResponse(response);
}

const parseResponse = async (response) => {
  const ok = response.ok;
  let json = ok ? await response.json() : null;
  return {
    status: response.status,
    ok,
    json
  };
}
