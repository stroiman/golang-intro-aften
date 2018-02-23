export const get = async (url, body) => {
  const response = await fetch(url, { body });
  const ok = response.ok;
  let json = ok ? await response.json() : null;
  return {
    status: response.status,
    ok,
    json
  };
};
