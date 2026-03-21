// main.ts

// Call this function to fetch boards from the backend and render them
function fetchBoards(): void {
  fetch("/boards")
    .then(async (response) => {
      if (!response.ok) throw new Error(await response.text());
      return response.json();
    })
    .then((boards: { id: string; name: string }[]) => {
      const ul = document.getElementById("boards-list");
      if (!ul) return;
      ul.innerHTML = "";
      boards.forEach((board) => {
        const li = document.createElement("li");
        li.textContent = board.name;
        ul.appendChild(li);
      });
    })
    .catch((error) => {
      alert("Error loading boards: " + error);
    });
}

// Attach to window for inline onclick, or set up an event listener:
(window as any).fetchBoards = fetchBoards;

// Optional: Fetch boards on page load
// window.addEventListener('DOMContentLoaded', fetchBoards);
