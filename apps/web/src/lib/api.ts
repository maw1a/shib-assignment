// * INFO: Load NEXT_PUBLIC_BASE_API_URL from .env file
const BASE_URL = process.env.NEXT_PUBLIC_BASE_API_URL;

export async function fetchTokenByOwner(owner: string | undefined) {
  if (!owner) return null;

  const res = await fetch(BASE_URL + "/nft?owner=" + owner);
  const json = await res.json();

  return json as {
    message: string;
    symbol: string;
    token_id: number;
    owner: string;
  };
}

export async function mintToken(to: string | undefined) {
  if (!to) return undefined;

  const headers = new Headers();
  headers.set("Content-Type", "application/json");

  const res = await fetch(BASE_URL + "/nft/mint", {
    method: "POST",
    headers,
    body: JSON.stringify({ to }),
  });

  const json: { message: string } = await res.json();

  return { ok: res.ok, message: json.message };
}
