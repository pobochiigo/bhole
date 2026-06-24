import './style.css';
import { createClient, createConnectTransport, Launch, Agency, SpaceStation, CelestialBody } from 'bhole-client';

// Define the API base URLs
const LOCAL_RPC_URL = 'http://localhost:8080';
const FALLBACK_REST_URL = 'https://lldev.thespacedevs.com/2.3.0';

// Setup ConnectRPC Transport and Clients
const transport = createConnectTransport({
  baseUrl: LOCAL_RPC_URL,
});

const launchClient = createClient(Launch.LaunchService, transport);
const agencyClient = createClient(Agency.AgencyService, transport);
const stationClient = createClient(SpaceStation.SpaceStationService, transport);
const bodyClient = createClient(CelestialBody.CelestialBodyService, transport);

// State tracking
let activeTab = 'launches';
let searchQuery = '';
let currentSource: 'connectrpc' | 'rest' | 'stubs' = 'connectrpc';
let loadedData: any[] = [];

// Curated high-resolution fallback space images
const IMG_LAUNCH = 'https://images.unsplash.com/photo-1541185933-ef5d8ed016c2?auto=format&fit=crop&w=800&q=80';
const IMG_AGENCY = 'https://images.unsplash.com/photo-1451187580459-43490279c0fa?auto=format&fit=crop&w=800&q=80';
const IMG_STATION = 'https://images.unsplash.com/photo-1446776811953-b23d57bd21aa?auto=format&fit=crop&w=800&q=80';
const IMG_BODY = 'https://images.unsplash.com/photo-1614730321146-b6fa6a46bcb4?auto=format&fit=crop&w=800&q=80';

// Curated offline fallback mock data (Stubs)
const OFFLINE_STUBS: Record<string, any[]> = {
  launches: [
    {
      id: 'l-1',
      name: 'Falcon Heavy | Psyche Mission',
      net: new Date(Date.now() + 86400000 * 5).toISOString(), // 5 days from now
      status: { name: 'Go for Launch', abbrev: 'Go', description: 'Systems are nominal. Ready for propellant loading.' },
      launchServiceProvider: { name: 'SpaceX', abbrev: 'SPX', featured: true, foundingYear: 2002 },
      image: { imageUrl: IMG_LAUNCH },
      mission: { name: 'Psyche Asteroid Survey', description: 'Psyche is a journey to a unique metal-rich asteroid orbiting the Sun between Mars and Jupiter. The mission will explore the origin of planetary cores.' },
      pad: { name: 'LC-39A', location: { name: 'Kennedy Space Center, Florida, USA' } }
    },
    {
      id: 'l-2',
      name: 'Artemis II | Crewed Lunar Flyby',
      net: new Date(Date.now() + 86400000 * 45).toISOString(), // 45 days from now
      status: { name: 'Scheduled', abbrev: 'TBD', description: 'First crewed mission of the Space Launch System (SLS) and the Orion spacecraft.' },
      launchServiceProvider: { name: 'NASA', abbrev: 'NASA', featured: true, foundingYear: 1958 },
      image: { imageUrl: 'https://images.unsplash.com/photo-1446776811953-b23d57bd21aa?auto=format&fit=crop&w=800&q=80' },
      mission: { name: 'Lunar Orbit Flyby', description: 'Artemis II will send four astronauts around the Moon to test the Orion spacecraft\'s life support systems in deep space.' },
      pad: { name: 'LC-39B', location: { name: 'Kennedy Space Center, Florida, USA' } }
    },
    {
      id: 'l-3',
      name: 'Ariane 6 | JUICE (Jupiter Icy Moons Explorer)',
      net: '2026-04-14T12:14:00Z',
      status: { name: 'Launch Successful', abbrev: 'Success', description: 'Spacecraft has separated and successfully entered heliocentric cruise trajectory.' },
      launchServiceProvider: { name: 'Arianespace', abbrev: 'ARI', featured: false, foundingYear: 1980 },
      image: { imageUrl: 'https://images.unsplash.com/photo-1517976487492-5750f3195933?auto=format&fit=crop&w=800&q=80' },
      mission: { name: 'JUICE', description: 'JUICE will make detailed observations of the giant gaseous planet Jupiter and its three large ocean-bearing moons: Ganymede, Callisto, and Europa.' },
      pad: { name: 'ELA-4', location: { name: 'Guiana Space Centre, Kourou, French Guiana' } }
    },
    {
      id: 'l-4',
      name: 'Electron | CAPSTONE Flight',
      net: '2026-06-28T09:55:00Z',
      status: { name: 'Launch Successful', abbrev: 'Success', description: 'Payload deployed into Lunar Near-Rectilinear Halo Orbit (NRHO).' },
      launchServiceProvider: { name: 'Rocket Lab', abbrev: 'RL', featured: true, foundingYear: 2006 },
      image: { imageUrl: 'https://images.unsplash.com/photo-1541185933-ef5d8ed016c2?auto=format&fit=crop&w=800&q=80' },
      mission: { name: 'CAPSTONE', description: 'Cislunar Autonomous Positioning System Technology Operations and Navigation Experiment. Designed to test the stability of the planned Gateway lunar orbit.' },
      pad: { name: 'LC-1A', location: { name: 'Mahia Peninsula, New Zealand' } }
    }
  ],
  agencies: [
    {
      id: 1,
      name: 'National Aeronautics and Space Administration',
      abbrev: 'NASA',
      typeVal: { name: 'Government' },
      country: [{ name: 'United States of America', alpha2Code: 'US' }],
      foundingYear: 1958,
      administrator: 'Bill Nelson',
      description: 'NASA is an independent agency of the US federal government responsible for the civil space program, aeronautics research, and space research.',
      image: { imageUrl: IMG_AGENCY },
      totalLaunchCount: 412,
      successfulLaunches: 385,
      failedLaunches: 27
    },
    {
      id: 2,
      name: 'Space Exploration Technologies Corp.',
      abbrev: 'SpaceX',
      typeVal: { name: 'Commercial' },
      country: [{ name: 'United States of America', alpha2Code: 'US' }],
      foundingYear: 2002,
      administrator: 'Elon Musk',
      description: 'SpaceX is an American aerospace manufacturer and space transportation services company headquartered in Hawthorne, California. It was founded in 2002 with the goal of reducing space transportation costs to enable the colonization of Mars.',
      image: { imageUrl: 'https://images.unsplash.com/photo-1516849841032-87cbac4d88f7?auto=format&fit=crop&w=800&q=80' },
      totalLaunchCount: 360,
      successfulLaunches: 358,
      failedLaunches: 2
    },
    {
      id: 3,
      name: 'European Space Agency',
      abbrev: 'ESA',
      typeVal: { name: 'Multinational' },
      country: [{ name: 'Europe', alpha2Code: 'EU' }],
      foundingYear: 1975,
      administrator: 'Josef Aschbacher',
      description: 'ESA is an intergovernmental organisation of 22 member states dedicated to the exploration of space.',
      image: { imageUrl: 'https://images.unsplash.com/photo-1446776811953-b23d57bd21aa?auto=format&fit=crop&w=800&q=80' },
      totalLaunchCount: 120,
      successfulLaunches: 114,
      failedLaunches: 6
    }
  ],
  stations: [
    {
      id: 1,
      name: 'International Space Station',
      status: { name: 'Active / Operational' },
      type: { name: 'Multinational' },
      founded: '1998-11-20',
      orbit: 'Low Earth Orbit (LEO)',
      description: 'The International Space Station is a modular space station in low Earth orbit. It is a multinational collaborative project between five participating space agencies: NASA, Roscosmos, JAXA, ESA, and CSA.',
      image: { imageUrl: IMG_STATION }
    },
    {
      id: 2,
      name: 'Tiangong Space Station',
      status: { name: 'Active / Operational' },
      type: { name: 'Government' },
      founded: '2021-04-29',
      orbit: 'Low Earth Orbit (LEO)',
      description: 'The Tiangong Space Station is a space station placed in Low Earth Orbit between 340 and 450 km above the Earth. It is operated by China Manned Space Agency.',
      image: { imageUrl: 'https://images.unsplash.com/photo-1451187580459-43490279c0fa?auto=format&fit=crop&w=800&q=80' }
    }
  ],
  bodies: [
    {
      id: 1,
      name: 'Earth',
      type: { name: 'Planet' },
      mass: 5.972e24,
      gravity: 9.807,
      atmosphere: true,
      diameter: 12742,
      description: 'Earth is the third planet from the Sun and the only astronomical object known to harbor life. About 29.2% of Earth\'s surface is land consisting of continents and islands.',
      image: { imageUrl: IMG_BODY }
    },
    {
      id: 2,
      name: 'Moon',
      type: { name: 'Natural Satellite' },
      mass: 7.342e22,
      gravity: 1.62,
      atmosphere: false,
      diameter: 3474,
      description: 'The Moon is Earth\'s only natural satellite. It is the fifth-largest satellite in the Solar System and the largest and most massive relative to its parent planet.',
      image: { imageUrl: 'https://images.unsplash.com/photo-1522030299830-16b8d3d049fe?auto=format&fit=crop&w=800&q=80' }
    },
    {
      id: 3,
      name: 'Mars',
      type: { name: 'Planet' },
      mass: 6.39e23,
      gravity: 3.720,
      atmosphere: true,
      diameter: 6779,
      description: 'Mars is the fourth planet from the Sun and the second-smallest planet in the Solar System, being larger than only Mercury. Mars carries the name of the Roman god of war.',
      image: { imageUrl: 'https://images.unsplash.com/photo-1614728894747-a83421e2b9c9?auto=format&fit=crop&w=800&q=80' }
    }
  ]
};

// Simple utility to recursively convert snake_case object keys to camelCase
function camelCaseKeys(obj: any): any {
  if (Array.isArray(obj)) {
    return obj.map(v => camelCaseKeys(v));
  } else if (obj !== null && obj !== undefined && obj.constructor === Object) {
    return Object.keys(obj).reduce((result, key) => {
      const camelKey = key.replace(/_([a-z0-9])/g, (_, g) => g.toUpperCase());
      result[camelKey] = camelCaseKeys(obj[key]);
      return result;
    }, {} as any);
  }
  return obj;
}

// Update connection status badges on UI
function updateConnectionBadges(rpc: 'online' | 'offline' | 'fetching', rest: 'online' | 'offline' | 'fetching', modeText: string) {
  const rpcDot = document.querySelector('#connect-rpc-status .indicator-dot');
  const restDot = document.querySelector('#rest-api-status .indicator-dot');
  const modeBadge = document.getElementById('app-mode-badge');

  if (rpcDot) {
    rpcDot.className = `indicator-dot ${rpc}`;
  }
  if (restDot) {
    restDot.className = `indicator-dot ${rest}`;
  }
  if (modeBadge) {
    modeBadge.textContent = modeText;
    if (modeText.includes('STUBS')) {
      modeBadge.style.background = 'linear-gradient(135deg, var(--primary), var(--secondary))';
    } else if (modeText.includes('REST')) {
      modeBadge.style.background = 'linear-gradient(135deg, var(--cyan), var(--primary))';
    } else {
      modeBadge.style.background = 'linear-gradient(135deg, #10b981, var(--cyan))';
    }
  }

  // Display status banner if using offline cache
  const banner = document.getElementById('status-banner');
  if (banner) {
    if (modeText.includes('STUBS')) {
      banner.textContent = '⚠️ System Alert: Unable to connect to local ConnectRPC or Remote REST API. Operating in Offline Cache Mode.';
      banner.classList.remove('hidden');
    } else if (modeText.includes('REST')) {
      banner.textContent = 'ℹ️ System Status: Local ConnectRPC server offline. Operating in Remote REST API Fallback Mode.';
      banner.classList.remove('hidden');
    } else {
      banner.classList.add('hidden');
    }
  }
}

// Fetch Data with full three-tiered fallback logic
async function fetchData(tab: string) {
  const loader = document.getElementById('loader');
  const deckGrid = document.getElementById('deck-grid');
  
  if (loader) loader.classList.remove('hidden');
  if (deckGrid) deckGrid.innerHTML = '';
  
  // Update state badges to "fetching"
  updateConnectionBadges(
    currentSource === 'connectrpc' ? 'fetching' : 'offline',
    currentSource === 'rest' ? 'fetching' : 'offline',
    'SYNCING...'
  );

  // --- TIER 1: Try ConnectRPC (Local) ---
  try {
    let results: any[] = [];
    if (tab === 'launches') {
      const resp = await launchClient.listLaunches({ limit: 15 });
      results = resp.results;
    } else if (tab === 'agencies') {
      const resp = await agencyClient.listAgencies({ limit: 15 });
      results = resp.results;
    } else if (tab === 'stations') {
      const resp = await stationClient.listSpaceStations({ limit: 15 });
      results = resp.results;
    } else if (tab === 'bodies') {
      const resp = await bodyClient.listCelestialBodies({ limit: 15 });
      results = resp.results;
    }
    
    currentSource = 'connectrpc';
    loadedData = results;
    updateConnectionBadges('online', 'offline', 'LOCAL CONNECTRPC');
    renderGrid();
    if (loader) loader.classList.add('hidden');
    return;
  } catch (rpcErr) {
    console.warn('ConnectRPC failed. Falling back to REST API.', rpcErr);
  }

  // --- TIER 2: Try public REST API (Remote) ---
  // Note: Celestial Bodies is a custom endpoint, so we bypass REST directly for bodies.
  if (tab !== 'bodies') {
    try {
      let endpoint = '';
      if (tab === 'launches') {
        endpoint = `${FALLBACK_REST_URL}/launches/upcoming/?limit=15`;
      } else if (tab === 'agencies') {
        endpoint = `${FALLBACK_REST_URL}/agencies/?limit=15`;
      } else if (tab === 'stations') {
        endpoint = `${FALLBACK_REST_URL}/space_stations/?limit=15`;
      }

      const response = await fetch(endpoint);
      if (!response.ok) throw new Error(`HTTP error ${response.status}`);
      const data = await response.json();
      
      // Convert keys from snake_case to camelCase to conform with Protobuf types
      const normalizedResults = camelCaseKeys(data.results || []);
      
      currentSource = 'rest';
      loadedData = normalizedResults;
      updateConnectionBadges('offline', 'online', 'REMOTE REST FALLBACK');
      renderGrid();
      if (loader) loader.classList.add('hidden');
      return;
    } catch (restErr) {
      console.warn('Fallback REST API failed. Falling back to Offline Cache (Stubs).', restErr);
    }
  }

  // --- TIER 3: Fall back to local stubs ---
  currentSource = 'stubs';
  loadedData = OFFLINE_STUBS[tab] || [];
  updateConnectionBadges('offline', 'offline', 'OFFLINE CACHE STUBS');
  renderGrid();
  if (loader) loader.classList.add('hidden');
}

// Filter and render the active items in the grid
function renderGrid() {
  const grid = document.getElementById('deck-grid');
  const totalCountEl = document.getElementById('total-count');
  if (!grid) return;

  const query = searchQuery.toLowerCase().trim();
  const filtered = loadedData.filter((item: any) => {
    if (!query) return true;
    const name = (item.name || '').toLowerCase();
    const desc = (item.description || item.mission?.description || '').toLowerCase();
    const abbrev = (item.abbrev || item.launchServiceProvider?.abbrev || '').toLowerCase();
    return name.includes(query) || desc.includes(query) || abbrev.includes(query);
  });

  if (totalCountEl) {
    totalCountEl.textContent = filtered.length.toString();
  }

  if (filtered.length === 0) {
    grid.innerHTML = `
      <div class="loader-container" style="grid-column: 1 / -1; padding: 60px 0;">
        <svg viewBox="0 0 24 24" width="48" height="48" stroke="var(--text-muted)" stroke-width="1.5" fill="none"><circle cx="12" cy="12" r="10"></circle><line x1="8" y1="12" x2="16" y2="12"></line></svg>
        <p style="margin-top: 12px; color: var(--text-secondary);">No telemetry matching filters found in current cache.</p>
      </div>
    `;
    return;
  }

  grid.innerHTML = filtered.map((item: any) => createCardHTML(item)).join('');

  // Attach event listeners to card clicks
  grid.querySelectorAll('.card').forEach(card => {
    card.addEventListener('click', () => {
      const id = card.getAttribute('data-id');
      const selected = filtered.find(x => String(x.id) === String(id));
      if (selected) {
        showModal(selected);
      }
    });
  });
}

// Generate the card element HTML
function createCardHTML(item: any): string {
  if (activeTab === 'launches') {
    const provider = item.launchServiceProvider?.name || 'Unknown Provider';
    const locationName = item.pad?.location?.name || 'Unknown Location';
    const statusText = item.status?.name || 'Scheduled';
    const isSuccess = statusText.toLowerCase().includes('success');
    const isUpcoming = statusText.toLowerCase().includes('go') || statusText.toLowerCase().includes('scheduled');
    
    let badgeClass = 'tag-failed';
    if (isSuccess) badgeClass = 'tag-success';
    else if (isUpcoming) badgeClass = 'tag-upcoming';

    const launchDate = new Date(item.net);
    const dateFormatted = launchDate.toLocaleDateString(undefined, {
      month: 'short', day: 'numeric', year: 'numeric', hour: '2-digit', minute: '2-digit'
    });

    const isLive = item.webcastLive ? '<span class="card-tag tag-upcoming" style="animation: dotBlink 1s infinite alternate;">LIVE STREAM</span>' : '';

    return `
      <article class="card" data-id="${item.id}">
        <div class="card-header">
          <span class="card-tag ${badgeClass}">${statusText}</span>
          ${isLive}
        </div>
        <div class="card-image-wrap">
          <img src="${item.image?.imageUrl || IMG_LAUNCH}" class="card-img" alt="${item.name}" loading="lazy" onerror="this.src='${IMG_LAUNCH}'" />
        </div>
        <h3 class="card-title">${item.name}</h3>
        <div class="card-subtitle">${provider} • ${locationName}</div>
        <div class="card-body">
          <p>${item.mission?.description || item.status?.description || 'No launch description available currently.'}</p>
        </div>
        <div class="card-footer">
          <span>${dateFormatted}</span>
          <span class="card-action">View Telemetry &rarr;</span>
        </div>
      </article>
    `;
  }
  
  if (activeTab === 'agencies') {
    const country = item.country?.[0]?.name || 'Global';
    const founding = item.foundingYear || 'N/A';
    const launches = item.totalLaunchCount || 0;
    
    return `
      <article class="card" data-id="${item.id}">
        <div class="card-header">
          <span class="card-tag tag-agency">${item.typeVal?.name || 'Space Agency'}</span>
        </div>
        <div class="card-image-wrap">
          <img src="${item.image?.imageUrl || IMG_AGENCY}" class="card-img" alt="${item.name}" loading="lazy" onerror="this.src='${IMG_AGENCY}'" />
        </div>
        <h3 class="card-title">${item.name} (${item.abbrev || 'N/A'})</h3>
        <div class="card-subtitle">Founded: ${founding} | ${country}</div>
        <div class="card-body">
          <p>${item.description || 'No agency biography or description is currently loaded.'}</p>
        </div>
        <div class="card-footer">
          <span>Launches: <strong>${launches}</strong></span>
          <span class="card-action">Agency Profile &rarr;</span>
        </div>
      </article>
    `;
  }
  
  if (activeTab === 'stations') {
    const orbit = item.orbit || 'Low Earth Orbit';
    const statusText = item.status?.name || 'Active';
    const foundedDate = item.founded ? new Date(item.founded).toLocaleDateString(undefined, { year: 'numeric', month: 'long' }) : 'N/A';
    
    return `
      <article class="card" data-id="${item.id}">
        <div class="card-header">
          <span class="card-tag tag-upcoming">${statusText}</span>
        </div>
        <div class="card-image-wrap">
          <img src="${item.image?.imageUrl || IMG_STATION}" class="card-img" alt="${item.name}" loading="lazy" onerror="this.src='${IMG_STATION}'" />
        </div>
        <h3 class="card-title">${item.name}</h3>
        <div class="card-subtitle">Orbit: ${orbit}</div>
        <div class="card-body">
          <div class="stat-row">
            <span class="stat-label">Inception Date</span>
            <span class="stat-value">${foundedDate}</span>
          </div>
          <div class="stat-row" style="margin-bottom: 12px;">
            <span class="stat-label">Operating Type</span>
            <span class="stat-value">${item.type?.name || 'Government'}</span>
          </div>
          <p>${item.description || 'Orbital platform tracking telemetry is operational.'}</p>
        </div>
        <div class="card-footer">
          <span>Station Data</span>
          <span class="card-action">Details &rarr;</span>
        </div>
      </article>
    `;
  }
  
  if (activeTab === 'bodies') {
    const gravity = item.gravity ? `${item.gravity.toFixed(2)} m/s²` : 'N/A';
    const atmos = item.atmosphere ? 'YES' : 'NO';
    const diameter = item.diameter ? `${item.diameter.toLocaleString()} km` : 'N/A';
    
    return `
      <article class="card" data-id="${item.id}">
        <div class="card-header">
          <span class="card-tag tag-body">${item.type?.name || 'Body'}</span>
        </div>
        <div class="card-image-wrap">
          <img src="${item.image?.imageUrl || IMG_BODY}" class="card-img" alt="${item.name}" loading="lazy" onerror="this.src='${IMG_BODY}'" />
        </div>
        <h3 class="card-title">${item.name}</h3>
        <div class="card-subtitle">Diameter: ${diameter}</div>
        <div class="card-body">
          <div class="stat-row">
            <span class="stat-label">Surface Gravity</span>
            <span class="stat-value">${gravity}</span>
          </div>
          <div class="stat-row" style="margin-bottom: 12px;">
            <span class="stat-label">Atmosphere</span>
            <span class="stat-value">${atmos}</span>
          </div>
          <p>${item.description || 'Celestial body catalog database details.'}</p>
        </div>
        <div class="card-footer">
          <span>Planetology Specs</span>
          <span class="card-action">Explore &rarr;</span>
        </div>
      </article>
    `;
  }

  return '';
}

// Display modal details
function showModal(item: any) {
  const modal = document.getElementById('details-modal');
  const catEl = document.getElementById('modal-category');
  const titleEl = document.getElementById('modal-title');
  const contentEl = document.getElementById('modal-content');
  
  if (!modal || !catEl || !titleEl || !contentEl) return;
  
  catEl.textContent = activeTab.toUpperCase();
  titleEl.textContent = item.name;
  
  let bodyHTML = '';
  
  if (activeTab === 'launches') {
    const launchDate = new Date(item.net);
    const dateFormatted = launchDate.toLocaleDateString(undefined, {
      weekday: 'long', month: 'long', day: 'numeric', year: 'numeric', hour: '2-digit', minute: '2-digit', timeZoneName: 'short'
    });
    
    const locationName = item.pad?.location?.name || 'N/A';
    const padName = item.pad?.name || 'N/A';
    const statusText = item.status?.name || 'Unknown';
    const provider = item.launchServiceProvider?.name || 'N/A';
    const desc = item.mission?.description || item.status?.description || 'No additional mission details available.';

    bodyHTML = `
      <img src="${item.image?.imageUrl || IMG_LAUNCH}" class="modal-hero-img" alt="${item.name}" onerror="this.src='${IMG_LAUNCH}'" />
      <div class="modal-section">
        <div class="modal-section-title">Flight Mission Profile</div>
        <p class="modal-desc">${desc}</p>
      </div>
      <div class="modal-section">
        <div class="modal-section-title">Telemetry Parameters</div>
        <div class="meta-grid">
          <div class="stat-row"><span class="stat-label">Launch Status</span><span class="stat-value" style="color: var(--cyan);">${statusText}</span></div>
          <div class="stat-row"><span class="stat-label">Operator</span><span class="stat-value">${provider}</span></div>
          <div class="stat-row"><span class="stat-label">Scheduled Lift-Off</span><span class="stat-value">${dateFormatted}</span></div>
          <div class="stat-row"><span class="stat-label">Launch Pad</span><span class="stat-value">${padName}</span></div>
          <div class="stat-row" style="grid-column: 1 / -1;"><span class="stat-label">Location Site</span><span class="stat-value">${locationName}</span></div>
        </div>
      </div>
    `;
  } else if (activeTab === 'agencies') {
    const admin = item.administrator || 'N/A';
    const founding = item.foundingYear || 'N/A';
    const country = item.country?.[0]?.name || 'N/A';
    const type = item.typeVal?.name || 'N/A';
    const launches = item.totalLaunchCount || 0;
    const successes = item.successfulLaunches || 0;
    const failures = item.failedLaunches || 0;
    
    bodyHTML = `
      <img src="${item.image?.imageUrl || IMG_AGENCY}" class="modal-hero-img" alt="${item.name}" onerror="this.src='${IMG_AGENCY}'" />
      <div class="modal-section">
        <div class="modal-section-title">Agency Overview</div>
        <p class="modal-desc">${item.description || 'No detailed biography loaded.'}</p>
      </div>
      <div class="modal-section">
        <div class="modal-section-title">Profile Parameters</div>
        <div class="meta-grid">
          <div class="stat-row"><span class="stat-label">Authority Type</span><span class="stat-value">${type}</span></div>
          <div class="stat-row"><span class="stat-label">HQ Region</span><span class="stat-value">${country}</span></div>
          <div class="stat-row"><span class="stat-label">Founding Year</span><span class="stat-value">${founding}</span></div>
          <div class="stat-row"><span class="stat-label">Current Admin</span><span class="stat-value">${admin}</span></div>
          <div class="stat-row"><span class="stat-label">Total Missions</span><span class="stat-value">${launches}</span></div>
          <div class="stat-row"><span class="stat-label">Success Rate</span><span class="stat-value" style="color: var(--status-success);">${launches > 0 ? ((successes / launches) * 100).toFixed(1) : 0}%</span></div>
          <div class="stat-row"><span class="stat-label">Successes</span><span class="stat-value">${successes}</span></div>
          <div class="stat-row"><span class="stat-label">Failures</span><span class="stat-value">${failures}</span></div>
        </div>
      </div>
    `;
  } else if (activeTab === 'stations') {
    const orbit = item.orbit || 'N/A';
    const founded = item.founded ? new Date(item.founded).toLocaleDateString(undefined, { year: 'numeric', month: 'long', day: 'numeric' }) : 'N/A';
    const type = item.type?.name || 'N/A';
    const status = item.status?.name || 'N/A';

    bodyHTML = `
      <img src="${item.image?.imageUrl || IMG_STATION}" class="modal-hero-img" alt="${item.name}" onerror="this.src='${IMG_STATION}'" />
      <div class="modal-section">
        <div class="modal-section-title">Orbital Facility Specs</div>
        <p class="modal-desc">${item.description || 'No station description available.'}</p>
      </div>
      <div class="modal-section">
        <div class="modal-section-title">Logistics Parameters</div>
        <div class="meta-grid">
          <div class="stat-row"><span class="stat-label">Facility Status</span><span class="stat-value" style="color: var(--cyan);">${status}</span></div>
          <div class="stat-row"><span class="stat-label">Station Type</span><span class="stat-value">${type}</span></div>
          <div class="stat-row"><span class="stat-label">Launch Inception</span><span class="stat-value">${founded}</span></div>
          <div class="stat-row"><span class="stat-label">Orbital Regime</span><span class="stat-value">${orbit}</span></div>
        </div>
      </div>
    `;
  } else if (activeTab === 'bodies') {
    const mass = item.mass ? `${item.mass.toExponential(3)} kg` : 'N/A';
    const gravity = item.gravity ? `${item.gravity.toFixed(3)} m/s²` : 'N/A';
    const diameter = item.diameter ? `${item.diameter.toLocaleString()} km` : 'N/A';
    const atmosphere = item.atmosphere ? 'Yes' : 'No';

    bodyHTML = `
      <img src="${item.image?.imageUrl || IMG_BODY}" class="modal-hero-img" alt="${item.name}" onerror="this.src='${IMG_BODY}'" />
      <div class="modal-section">
        <div class="modal-section-title">Planetological Data</div>
        <p class="modal-desc">${item.description || 'No planetological description currently loaded.'}</p>
      </div>
      <div class="modal-section">
        <div class="modal-section-title">Astrophysical Specifications</div>
        <div class="meta-grid">
          <div class="stat-row"><span class="stat-label">Planetology Class</span><span class="stat-value">${item.type?.name || 'N/A'}</span></div>
          <div class="stat-row"><span class="stat-label">Physical Diameter</span><span class="stat-value">${diameter}</span></div>
          <div class="stat-row"><span class="stat-label">Planetary Mass</span><span class="stat-value">${mass}</span></div>
          <div class="stat-row"><span class="stat-label">Surface Gravity</span><span class="stat-value">${gravity}</span></div>
          <div class="stat-row" style="grid-column: 1 / -1;"><span class="stat-label">Sustained Atmosphere</span><span class="stat-value">${atmosphere}</span></div>
        </div>
      </div>
    `;
  }
  
  contentEl.innerHTML = bodyHTML;
  modal.classList.remove('hidden');
}

// Close modal view
function closeModal() {
  const modal = document.getElementById('details-modal');
  if (modal) modal.classList.add('hidden');
}

// Global clock tick
function updateClock() {
  const clockEl = document.getElementById('utc-clock');
  if (clockEl) {
    const now = new Date();
    clockEl.textContent = now.toISOString().substring(11, 19);
  }
}

// Initialize Application UI Controls
document.addEventListener('DOMContentLoaded', () => {
  // Setup clock interval
  setInterval(updateClock, 1000);
  updateClock();

  // Setup Global Search filter
  const searchInput = document.getElementById('global-search') as HTMLInputElement;
  if (searchInput) {
    searchInput.addEventListener('input', (e) => {
      searchQuery = (e.target as HTMLInputElement).value;
      renderGrid();
    });
  }

  // Setup Navigation Tab Switcher
  const navItems = document.querySelectorAll('.nav-menu .nav-item');
  const tabTitle = document.getElementById('tab-title');
  const tabSubtitle = document.getElementById('tab-subtitle');

  navItems.forEach(item => {
    item.addEventListener('click', () => {
      const tab = item.getAttribute('data-tab');
      if (!tab || tab === activeTab) return;

      // Update active nav class
      navItems.forEach(x => x.classList.remove('active'));
      item.classList.add('active');

      activeTab = tab;
      
      // Update headings
      if (tabTitle && tabSubtitle) {
        if (tab === 'launches') {
          tabTitle.textContent = 'Space Launches';
          tabSubtitle.textContent = 'Real-time status monitor of upcoming and historical space operations';
        } else if (tab === 'agencies') {
          tabTitle.textContent = 'Space Agencies';
          tabSubtitle.textContent = 'Global governing bodies and commercial operators driving celestial transport';
        } else if (tab === 'stations') {
          tabTitle.textContent = 'Space Stations';
          tabSubtitle.textContent = 'Modular orbital installations orbiting high above the atmosphere';
        } else if (tab === 'bodies') {
          tabTitle.textContent = 'Celestial Bodies';
          tabSubtitle.textContent = 'A catalog of planetary targets, moons, and astronomical destinations';
        }
      }

      // Re-fetch tab content
      fetchData(tab);
    });
  });

  // Modal close handlers
  const closeBtn = document.getElementById('modal-close-btn');
  const modalOverlay = document.getElementById('details-modal');

  if (closeBtn) {
    closeBtn.addEventListener('click', closeModal);
  }
  if (modalOverlay) {
    modalOverlay.addEventListener('click', (e) => {
      if (e.target === modalOverlay) closeModal();
    });
  }

  // Escape key closes modal
  window.addEventListener('keydown', (e) => {
    if (e.key === 'Escape') closeModal();
  });

  // Initial fetch for the default tab
  fetchData('launches');
});
