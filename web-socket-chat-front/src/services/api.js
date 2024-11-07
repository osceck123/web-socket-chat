// src/api.js
const API_URL = "http://localhost:8080";

// Inicio de sesión
export async function login(username, password) {
  const response = await fetch(`${API_URL}/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ username, password })
  });

  if (response.ok) {
    return response.json();
  } else {
    throw new Error("Login failed");
  }
}

// Cierre de sesión
export async function logout() {
  await fetch(`${API_URL}/logout`, { method: "POST", credentials: "include" });
}

// Obtener acceso a la videollamada (requiere token)
export async function initiateVideoCall() {
  const response = await fetch(`${API_URL}/api/video`, {
    method: "GET",
    credentials: "include"
  });

  if (!response.ok) {
    throw new Error("Failed to initiate video call");
  }
}
