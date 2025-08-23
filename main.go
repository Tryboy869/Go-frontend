package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// ===== SVELTE PHILOSOPHY ABSORBED IN GO (ZERO EXTERNAL DEPS) =====

// ReactiveState : Go qui pense comme Svelte store - FROM SCRATCH
type ReactiveState struct {
	count      int
	doubled    int
	users      int
	lastUpdate string
	mu         sync.RWMutex
	channels   map[string]chan string // Go channels = Svelte reactivity
}

func NewReactiveState() *ReactiveState {
	rs := &ReactiveState{
		count:      0,
		doubled:    0,
		users:      0,
		lastUpdate: "Never",
		channels:   make(map[string]chan string),
	}
	
	// SVELTE PHILOSOPHY: Reactive propagation via Go channels
	rs.channels["updates"] = make(chan string, 100)
	
	// Go goroutine = Svelte reactive loop
	go rs.reactiveLoop()
	
	return rs
}

// SVELTE $: REACTIVE STATEMENTS - Go native implementation
func (rs *ReactiveState) reactiveLoop() {
	for update := range rs.channels["updates"] {
		// SVELTE THINKING: Auto-recalculate derived values
		rs.mu.Lock()
		rs.doubled = rs.count * 2        // $: doubled = count * 2
		rs.lastUpdate = getCurrentTime()  // $: lastUpdate = now()
		rs.mu.Unlock()
		
		// SVELTE PHILOSOPHY: Broadcast changes (like Svelte stores)
		rs.broadcastUpdate(update)
	}
}

func (rs *ReactiveState) broadcastUpdate(update string) {
	// Go channels = Svelte reactive notifications
	fmt.Printf("🔄 Reactive update: %s\n", update)
}

// SVELTE STORE METHODS - Go native implementation
func (rs *ReactiveState) SetCount(newCount int) {
	rs.mu.Lock()
	rs.count = newCount
	rs.mu.Unlock()
	
	// SVELTE PHILOSOPHY: Trigger reactivity
	select {
	case rs.channels["updates"] <- fmt.Sprintf("count=%d", newCount):
	default:
	}
}

func (rs *ReactiveState) GetCount() int {
	rs.mu.RLock()
	defer rs.mu.RUnlock()
	return rs.count
}

func (rs *ReactiveState) GetDoubled() int {
	rs.mu.RLock()
	defer rs.mu.RUnlock()
	return rs.doubled
}

func (rs *ReactiveState) AddUser() {
	rs.mu.Lock()
	rs.users++
	rs.mu.Unlock()
	
	select {
	case rs.channels["updates"] <- fmt.Sprintf("users=%d", rs.users):
	default:
	}
}

func (rs *ReactiveState) GetUsers() int {
	rs.mu.RLock()
	defer rs.mu.RUnlock()
	return rs.users
}

func (rs *ReactiveState) GetLastUpdate() string {
	rs.mu.RLock()
	defer rs.mu.RUnlock()
	return rs.lastUpdate
}

// ===== HTML GENERATION - GO PURE (Svelte-like components) =====

// SVELTE COMPONENT PHILOSOPHY: Single file component logic
func (rs *ReactiveState) generateHTML() string {
	// SVELTE THINKING: Template + Logic + Style in one place
	// Go string building = Svelte template compilation
	
	var html strings.Builder
	
	// HTML Structure (Svelte-like template)
	html.WriteString("<!DOCTYPE html>\n")
	html.WriteString("<html>\n<head>\n")
	html.WriteString("<title>Go-Svelte Philosophy</title>\n")
	html.WriteString("<meta charset=\"utf-8\">\n")
	html.WriteString("<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n")
	
	// SVELTE <style> - Go string = Svelte CSS
	html.WriteString("<style>\n")
	html.WriteString(rs.generateCSS())
	html.WriteString("</style>\n")
	html.WriteString("</head>\n<body>\n")
	
	// SVELTE {#each} logic - Go loops = Svelte iterations  
	html.WriteString("<div class=\"container\">\n")
	html.WriteString("<h1>🚀 Go Frontend with Svelte Philosophy</h1>\n")
	html.WriteString("<div class=\"badge\">Zero JS/CSS/HTML imports - Pure Go thinking like Svelte</div>\n")
	
	// SVELTE REACTIVE DISPLAY - Go string formatting = Svelte interpolation
	rs.mu.RLock()
	count := rs.count
	doubled := rs.doubled
	users := rs.users
	lastUpdate := rs.lastUpdate
	rs.mu.RUnlock()
	
	html.WriteString("<div class=\"stats\">\n")
	
	// Stat cards (Svelte component style)
	html.WriteString(rs.generateStatCard("Count", strconv.Itoa(count), "🔢"))
	html.WriteString(rs.generateStatCard("Doubled", strconv.Itoa(doubled), "✖️"))
	html.WriteString(rs.generateStatCard("Users", strconv.Itoa(users), "👥"))
	html.WriteString(rs.generateStatCard("Last Update", lastUpdate, "⏰"))
	
	html.WriteString("</div>\n")
	
	// SVELTE BUTTONS - Go forms = Svelte event handlers
	html.WriteString("<div class=\"controls\">\n")
	html.WriteString("<button onclick=\"fetch('/increment')\">➕ Increment</button>\n")
	html.WriteString("<button onclick=\"fetch('/decrement')\">➖ Decrement</button>\n")
	html.WriteString("<button onclick=\"fetch('/add-user')\">👤 Add User</button>\n")
	html.WriteString("<button onclick=\"fetch('/reset')\">🔄 Reset</button>\n")
	html.WriteString("</div>\n")
	
	// SVELTE REACTIVITY - Go = Auto-refresh like Svelte
	html.WriteString("<div class=\"auto-refresh\">\n")
	html.WriteString("<p>🔄 Auto-refreshing every 2 seconds (Svelte-like reactivity)</p>\n")
	html.WriteString("</div>\n")
	
	html.WriteString("</div>\n")
	
	// SVELTE <script> - Go JavaScript = Svelte client-side
	html.WriteString("<script>\n")
	html.WriteString(rs.generateJavaScript())
	html.WriteString("</script>\n")
	
	html.WriteString("</body>\n</html>")
	
	return html.String()
}

// SVELTE CSS - Go string building = Svelte style compilation
func (rs *ReactiveState) generateCSS() string {
	var css strings.Builder
	
	css.WriteString("body {\n")
	css.WriteString("  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;\n")
	css.WriteString("  max-width: 900px;\n")
	css.WriteString("  margin: 0 auto;\n")
	css.WriteString("  padding: 20px;\n")
	css.WriteString("  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);\n")
	css.WriteString("  min-height: 100vh;\n")
	css.WriteString("}\n")
	
	css.WriteString(".container {\n")
	css.WriteString("  background: white;\n")
	css.WriteString("  border-radius: 15px;\n")
	css.WriteString("  padding: 30px;\n")
	css.WriteString("  box-shadow: 0 20px 40px rgba(0,0,0,0.1);\n")
	css.WriteString("}\n")
	
	css.WriteString("h1 {\n")
	css.WriteString("  color: #333;\n")
	css.WriteString("  text-align: center;\n")
	css.WriteString("  margin-bottom: 20px;\n")
	css.WriteString("}\n")
	
	css.WriteString(".badge {\n")
	css.WriteString("  background: #4CAF50;\n")
	css.WriteString("  color: white;\n")
	css.WriteString("  padding: 8px 15px;\n")
	css.WriteString("  border-radius: 20px;\n")
	css.WriteString("  font-size: 13px;\n")
	css.WriteString("  text-align: center;\n")
	css.WriteString("  margin-bottom: 30px;\n")
	css.WriteString("  font-weight: bold;\n")
	css.WriteString("}\n")
	
	css.WriteString(".stats {\n")
	css.WriteString("  display: grid;\n")
	css.WriteString("  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));\n")
	css.WriteString("  gap: 20px;\n")
	css.WriteString("  margin: 30px 0;\n")
	css.WriteString("}\n")
	
	css.WriteString(".stat-card {\n")
	css.WriteString("  background: #f8f9fa;\n")
	css.WriteString("  padding: 20px;\n")
	css.WriteString("  border-radius: 10px;\n")
	css.WriteString("  border-left: 4px solid #667eea;\n")
	css.WriteString("  text-align: center;\n")
	css.WriteString("}\n")
	
	css.WriteString(".stat-emoji {\n")
	css.WriteString("  font-size: 24px;\n")
	css.WriteString("  margin-bottom: 10px;\n")
	css.WriteString("}\n")
	
	css.WriteString(".stat-title {\n")
	css.WriteString("  font-size: 14px;\n")
	css.WriteString("  color: #666;\n")
	css.WriteString("  margin-bottom: 5px;\n")
	css.WriteString("}\n")
	
	css.WriteString(".stat-value {\n")
	css.WriteString("  font-size: 32px;\n")
	css.WriteString("  font-weight: bold;\n")
	css.WriteString("  color: #333;\n")
	css.WriteString("}\n")
	
	css.WriteString(".controls {\n")
	css.WriteString("  display: flex;\n")
	css.WriteString("  gap: 15px;\n")
	css.WriteString("  justify-content: center;\n")
	css.WriteString("  flex-wrap: wrap;\n")
	css.WriteString("  margin: 30px 0;\n")
	css.WriteString("}\n")
	
	css.WriteString("button {\n")
	css.WriteString("  background: linear-gradient(45deg, #667eea, #764ba2);\n")
	css.WriteString("  color: white;\n")
	css.WriteString("  border: none;\n")
	css.WriteString("  padding: 12px 24px;\n")
	css.WriteString("  border-radius: 25px;\n")
	css.WriteString("  cursor: pointer;\n")
	css.WriteString("  font-size: 16px;\n")
	css.WriteString("  font-weight: bold;\n")
	css.WriteString("  transition: transform 0.2s;\n")
	css.WriteString("}\n")
	
	css.WriteString("button:hover {\n")
	css.WriteString("  transform: translateY(-2px);\n")
	css.WriteString("}\n")
	
	css.WriteString(".auto-refresh {\n")
	css.WriteString("  text-align: center;\n")
	css.WriteString("  margin-top: 30px;\n")
	css.WriteString("  color: #666;\n")
	css.WriteString("  font-style: italic;\n")
	css.WriteString("}\n")
	
	return css.String()
}

// SVELTE COMPONENT HELPER - Go function = Svelte component
func (rs *ReactiveState) generateStatCard(title, value, emoji string) string {
	var card strings.Builder
	
	card.WriteString("<div class=\"stat-card\">\n")
	card.WriteString(fmt.Sprintf("  <div class=\"stat-emoji\">%s</div>\n", emoji))
	card.WriteString(fmt.Sprintf("  <div class=\"stat-title\">%s</div>\n", title))
	card.WriteString(fmt.Sprintf("  <div class=\"stat-value\">%s</div>\n", value))
	card.WriteString("</div>\n")
	
	return card.String()
}

// SVELTE CLIENT REACTIVITY - Go JavaScript = Svelte client updates
func (rs *ReactiveState) generateJavaScript() string {
	var js strings.Builder
	
	js.WriteString("// SVELTE PHILOSOPHY: Auto-refresh reactivity\n")
	js.WriteString("setInterval(() => {\n")
	js.WriteString("  fetch('/api/state')\n")
	js.WriteString("    .then(response => response.json())\n")
	js.WriteString("    .then(data => {\n")
	js.WriteString("      // Go-generated JS = Svelte reactive updates\n")
	js.WriteString("      if (data.shouldRefresh) {\n")
	js.WriteString("        window.location.reload();\n")
	js.WriteString("      }\n")
	js.WriteString("    })\n")
	js.WriteString("    .catch(() => {});\n")
	js.WriteString("}, 2000);\n")
	
	js.WriteString("\n// SVELTE-style button interactions\n")
	js.WriteString("document.addEventListener('click', (e) => {\n")
	js.WriteString("  if (e.target.tagName === 'BUTTON') {\n")
	js.WriteString("    e.target.style.transform = 'scale(0.95)';\n")
	js.WriteString("    setTimeout(() => {\n")
	js.WriteString("      e.target.style.transform = '';\n")
	js.WriteString("      window.location.reload();\n")
	js.WriteString("    }, 100);\n")
	js.WriteString("  }\n")
	js.WriteString("});\n")
	
	return js.String()
}

// ===== UTILITY FUNCTIONS (Go Pure) =====

func getCurrentTime() string {
	return time.Now().Format("15:04:05")
}

// ===== GLOBAL STATE (Svelte Store Philosophy) =====

var globalReactiveState = NewReactiveState()

// ===== HTTP HANDLERS (Go + Svelte Philosophy) =====

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// SVELTE THINKING: Component renders with current state
	html := globalReactiveState.generateHTML()
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
	
	fmt.Printf("🏠 Home rendered - Count: %d, Users: %d\n", 
		globalReactiveState.GetCount(), globalReactiveState.GetUsers())
}

func incrementHandler(w http.ResponseWriter, r *http.Request) {
	// SVELTE PHILOSOPHY: Action triggers reactive update
	newCount := globalReactiveState.GetCount() + 1
	globalReactiveState.SetCount(newCount)
	
	fmt.Printf("➕ Increment - New count: %d (doubled: %d)\n", 
		newCount, globalReactiveState.GetDoubled())
	
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func decrementHandler(w http.ResponseWriter, r *http.Request) {
	newCount := globalReactiveState.GetCount() - 1
	if newCount < 0 {
		newCount = 0
	}
	globalReactiveState.SetCount(newCount)
	
	fmt.Printf("➖ Decrement - New count: %d\n", newCount)
	
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	globalReactiveState.AddUser()
	
	fmt.Printf("👤 User added - Total users: %d\n", globalReactiveState.GetUsers())
	
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	globalReactiveState.SetCount(0)
	globalReactiveState.mu.Lock()
	globalReactiveState.users = 0
	globalReactiveState.mu.Unlock()
	
	fmt.Printf("🔄 Reset - All counters cleared\n")
	
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func stateAPIHandler(w http.ResponseWriter, r *http.Request) {
	// SVELTE THINKING: API for reactive updates
	state := map[string]interface{}{
		"count":        globalReactiveState.GetCount(),
		"doubled":      globalReactiveState.GetDoubled(),
		"users":        globalReactiveState.GetUsers(),
		"lastUpdate":   globalReactiveState.GetLastUpdate(),
		"shouldRefresh": false,
	}
	
	w.Header().Set("Content-Type", "application/json")
	
	// Go JSON encoding = Svelte data serialization
	jsonData := fmt.Sprintf(`{
		"count": %d,
		"doubled": %d,
		"users": %d,
		"lastUpdate": "%s",
		"shouldRefresh": false
	}`, state["count"], state["doubled"], state["users"], state["lastUpdate"])
	
	w.Write([]byte(jsonData))
}

// ===== MAIN SERVER (Go Simplicity + Svelte Reactivity) =====

func main() {
	fmt.Println("🚀 Go-Svelte Philosophy Server Starting...")
	fmt.Println("📱 NEXUS AXION: Go reasoning like Svelte for frontend")
	fmt.Println("✨ Zero external deps - Pure philosophy absorption")
	
	// SVELTE PHILOSOPHY: Component-based routing
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/increment", incrementHandler)
	http.HandleFunc("/decrement", decrementHandler)
	http.HandleFunc("/add-user", addUserHandler)
	http.HandleFunc("/reset", resetHandler)
	http.HandleFunc("/api/state", stateAPIHandler)
	
	// Go channels for background reactivity (Svelte-style)
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		
		for range ticker.C {
			// Background reactive updates
			fmt.Printf("🔄 Reactive heartbeat - State: count=%d, users=%d\n", 
				globalReactiveState.GetCount(), globalReactiveState.GetUsers())
		}
	}()
	
	// Port configuration for deployment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	fmt.Printf("🌐 Server listening on port %s\n", port)
	fmt.Printf("🔗 Open: http://localhost:%s\n", port)
	fmt.Println("💡 Philosophy: Go thinking like Svelte - Zero simulation!")
	
	log.Fatal(http.ListenAndServe(":"+port, nil))
}