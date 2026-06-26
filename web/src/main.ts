import './style.css';
import { createClient, createConnectTransport, Launch, Agency, SpaceStation, CelestialBody } from '@pobochiigo/bhole-client';

// Define the API base URLs
const urlParams = new URLSearchParams(window.location.search);
const LOCAL_RPC_URL = urlParams.get('api') || import.meta.env.VITE_API_URL || 'http://localhost:8080';
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

// Telemetry State variables for SpaceX Live Flight Dashboard
let telemetryInterval: any = null;
let telemetryTime = 0; // in seconds
let telemetrySpeed = 0; // km/h
let telemetryAlt = 0; // km
let telemetryDist = 0; // km downrange
let telemetryStage1Fuel = 100; // %
let telemetryStage2Fuel = 100; // %
let telemetryAborted = false;
let telemetryThrottle = 100; // %
let telemetryChartPoints: { x: number, y: number }[] = [];

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

// Deep clean mentions of Russia/Roscosmos/Soviet to sanitize multinational entries
function sanitizeRussiaMentions(obj: any): any {
  if (obj === null || obj === undefined) return obj;
  if (typeof obj === 'string') {
    return obj
      .replace(/roscosmos/gi, 'partner space agencies')
      .replace(/russian federal space agency/gi, 'partner space agency')
      .replace(/russian/gi, 'partner')
      .replace(/russia/gi, 'partner region');
  }
  if (Array.isArray(obj)) {
    return obj.map(item => sanitizeRussiaMentions(item));
  }
  if (typeof obj === 'object') {
    const copy = { ...obj };
    for (const key of Object.keys(copy)) {
      copy[key] = sanitizeRussiaMentions(copy[key]);
    }
    return copy;
  }
  return obj;
}

// Filter out Russian-operated, Russian-located, or Russian-owned entries entirely
function filterRussiaRelated(data: any[], tab: string): any[] {
  return data
    .filter((item: any) => {
      if (tab === 'agencies') {
        const countryCodes = (item.country || []).map((c: any) => (c.alpha2Code || c.alpha_2_code || '').toUpperCase());
        if (countryCodes.includes('RU') || countryCodes.includes('RUS')) return false;
        const name = (item.name || '').toLowerCase();
        const abbrev = (item.abbrev || '').toLowerCase();
        if (name.includes('roscosmos') || name.includes('russia') || name.includes('soviet') || abbrev.includes('roscosmos')) return false;
      }
      
      if (tab === 'launches') {
        const provider = item.launchServiceProvider || item.lsp;
        if (provider) {
          const name = (provider.name || '').toLowerCase();
          const abbrev = (provider.abbrev || '').toLowerCase();
          if (name.includes('roscosmos') || name.includes('russia') || name.includes('soviet') || abbrev.includes('roscosmos')) return false;
          const countryCodes = (provider.country || []).map((c: any) => (c.alpha2Code || c.alpha_2_code || '').toUpperCase());
          if (countryCodes.includes('RU') || countryCodes.includes('RUS')) return false;
        }
        const pad = item.pad;
        if (pad && pad.location) {
          const locName = (pad.location.name || '').toLowerCase();
          if (locName.includes('russia') || locName.includes('plesetsk') || locName.includes('vostochny') || locName.includes('baikonur')) return false;
        }
      }

      if (tab === 'stations') {
        const name = (item.name || '').toLowerCase();
        if (name.includes('mir') || name.includes('salyut') || name.includes('almaz')) return false;
        const owner = (item.owner || item.spaceAgency || item.agency || '').toLowerCase();
        if (owner.includes('roscosmos') || owner.includes('russia')) return false;
      }

      return true;
    })
    .map((item: any) => sanitizeRussiaMentions(item));
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
    loadedData = filterRussiaRelated(results, tab);
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
      loadedData = filterRussiaRelated(normalizedResults, tab);
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
  loadedData = filterRussiaRelated(OFFLINE_STUBS[tab] || [], tab);
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
  
  const card = modal.querySelector('.modal-card');
  let bodyHTML = '';
  
  if (activeTab === 'launches') {
    card?.classList.add('modal-spacex');
    bodyHTML = `
      <div class="abort-alert-overlay" id="abort-overlay"></div>
      <div class="spacex-console">
        <!-- Milestone Stepper -->
        <div class="milestone-stepper">
          <div class="milestone-step active" id="step-countdown">
            <div class="step-marker">1</div>
            <div class="step-info">
              <span class="step-title">Countdown</span>
              <span class="step-time" id="step-countdown-time">T-10.0s</span>
            </div>
          </div>
          <div class="milestone-step" id="step-liftoff">
            <div class="step-marker">2</div>
            <div class="step-info">
              <span class="step-title">Liftoff</span>
              <span class="step-time">T+0.0s</span>
            </div>
          </div>
          <div class="milestone-step" id="step-maxq">
            <div class="step-marker">3</div>
            <div class="step-info">
              <span class="step-title">Max Q</span>
              <span class="step-time">T+60.0s</span>
            </div>
          </div>
          <div class="milestone-step" id="step-meco">
            <div class="step-marker">4</div>
            <div class="step-info">
              <span class="step-title">MECO</span>
              <span class="step-time">T+150.0s</span>
            </div>
          </div>
          <div class="milestone-step" id="step-sep">
            <div class="step-marker">5</div>
            <div class="step-info">
              <span class="step-title">Stage Sep</span>
              <span class="step-time">T+160.0s</span>
            </div>
          </div>
          <div class="milestone-step" id="step-orbit">
            <div class="step-marker">6</div>
            <div class="step-info">
              <span class="step-title">Orbit Insertion</span>
              <span class="step-time">T+500.0s</span>
            </div>
          </div>
        </div>

        <!-- Telemetry Panel -->
        <div class="telemetry-panel">
          <div class="telemetry-deck">
            <div class="telemetry-card-spacex">
              <div class="telemetry-label-spacex">Speed</div>
              <div class="telemetry-value-spacex" id="telemetry-speed">0 <span class="telemetry-unit-spacex">km/h</span></div>
            </div>
            <div class="telemetry-card-spacex alt">
              <div class="telemetry-label-spacex">Altitude</div>
              <div class="telemetry-value-spacex" id="telemetry-alt">0.0 <span class="telemetry-unit-spacex">km</span></div>
            </div>
            <div class="telemetry-card-spacex">
              <div class="telemetry-label-spacex">Downrange</div>
              <div class="telemetry-value-spacex" id="telemetry-dist">0.0 <span class="telemetry-unit-spacex">km</span></div>
            </div>
          </div>

          <!-- Fuel progress -->
          <div class="fuel-tank-wrap">
            <div class="tank-label">
              <span>Stage 1 Propellant</span>
              <span id="fuel-stage1-val">100%</span>
            </div>
            <div class="tank-progress">
              <div class="tank-fill" id="fuel-stage1-fill" style="width: 100%;"></div>
            </div>
            <div style="margin-top: 10px;" class="tank-label">
              <span>Stage 2 Propellant</span>
              <span id="fuel-stage2-val">100%</span>
            </div>
            <div class="tank-progress">
              <div class="tank-fill stage2" id="fuel-stage2-fill" style="width: 100%;"></div>
            </div>
          </div>

          <!-- Canvas Chart -->
          <div class="chart-container-spacex">
            <canvas class="chart-canvas" id="telemetry-chart"></canvas>
          </div>
        </div>

        <!-- Right Control Panel -->
        <div class="control-section-spacex">
          <div class="flight-clock-panel">
            <span class="clock-title-spacex">Mission Clock</span>
            <span class="clock-value-spacex" id="telemetry-clock">T-10.0s</span>
          </div>

          <div class="override-panel-spacex">
            <div class="slider-value-wrap">
              <span class="override-label-spacex">Engine Thrust</span>
              <span class="slider-percentage" id="throttle-val">100%</span>
            </div>
            <input type="range" min="0" max="150" value="100" class="override-slider-spacex" id="throttle-override" />
          </div>

          <button class="abort-btn-spacex" id="abort-btn">Abort Mission</button>

          <div class="terminal-log-spacex" id="telemetry-terminal">
            <div class="log-line green">[SYSTEM] Telemetry simulation ready.</div>
            <div class="log-line cyan">[SYSTEM] Awaiting T-10 countdown...</div>
          </div>
        </div>
      </div>
    `;
    contentEl.innerHTML = bodyHTML;
    modal.classList.remove('hidden');
    
    // Start telemetry simulation
    startTelemetrySimulation();

  } else {
    card?.classList.remove('modal-spacex');
    if (activeTab === 'agencies') {
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
}

// Close modal view
function closeModal() {
  const modal = document.getElementById('details-modal');
  if (modal) {
    modal.classList.add('hidden');
    const card = modal.querySelector('.modal-card');
    card?.classList.remove('modal-spacex');
  }
  if (telemetryInterval) {
    clearInterval(telemetryInterval);
    telemetryInterval = null;
  }
}

function setActiveStep(stepId: string) {
  const steps = ['step-countdown', 'step-liftoff', 'step-maxq', 'step-meco', 'step-sep', 'step-orbit'];
  const activeIdx = steps.indexOf(stepId);
  if (activeIdx === -1) return;
  
  steps.forEach((id, idx) => {
    const el = document.getElementById(id);
    if (!el) return;
    if (idx < activeIdx) {
      el.classList.remove('active');
      el.classList.add('completed');
    } else if (idx === activeIdx) {
      el.classList.remove('completed');
      el.classList.add('active');
    } else {
      el.classList.remove('active', 'completed');
    }
  });
}

function addLogLine(text: string, color: 'cyan' | 'pink' | 'red' | 'green' | 'white' = 'white') {
  const terminal = document.getElementById('telemetry-terminal');
  if (!terminal) return;
  const line = document.createElement('div');
  line.className = `log-line ${color}`;
  
  let timeStr = '';
  if (telemetryTime < 0) {
    timeStr = `T${telemetryTime.toFixed(1)}s`;
  } else {
    const minutes = Math.floor(telemetryTime / 60);
    const seconds = Math.floor(telemetryTime % 60);
    timeStr = `T+${minutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}`;
  }
  line.textContent = `[${timeStr}] ${text}`;
  terminal.appendChild(line);
  terminal.scrollTop = terminal.scrollHeight;
}

function initiateAbort() {
  telemetryAborted = true;
  telemetryThrottle = 0;
  
  addLogLine('!!! EMERGENCY ABORT SEQUENCE ACTIVATED !!!', 'red');
  addLogLine('[SAFETY] Initiating Stage 1 Engine Cutoff (MECO).', 'red');
  addLogLine('[SAFETY] Commanding Stage separation & capsule escape thrusters.', 'pink');
  addLogLine('[SAFETY] Flight controls locked. Re-vectoring to ballistic descent.', 'red');

  const abortOverlay = document.getElementById('abort-overlay');
  if (abortOverlay) abortOverlay.classList.add('active');

  const abortBtn = document.getElementById('abort-btn');
  if (abortBtn) {
    abortBtn.textContent = 'ABORTED';
    abortBtn.classList.add('aborted');
  }

  const throttleSlider = document.getElementById('throttle-override') as HTMLInputElement;
  if (throttleSlider) {
    throttleSlider.disabled = true;
  }
}

function drawTelemetryChart() {
  const canvas = document.getElementById('telemetry-chart') as HTMLCanvasElement;
  if (!canvas) return;
  const ctx = canvas.getContext('2d');
  if (!ctx) return;

  const rect = canvas.getBoundingClientRect();
  const dpr = window.devicePixelRatio || 1;
  canvas.width = rect.width * dpr;
  canvas.height = rect.height * dpr;
  
  ctx.save();
  ctx.scale(dpr, dpr);

  const w = rect.width;
  const h = rect.height;

  // Clear background
  ctx.fillStyle = '#040308';
  ctx.fillRect(0, 0, w, h);

  // Draw Grid Lines (Grey dashed)
  ctx.strokeStyle = 'rgba(255, 255, 255, 0.05)';
  ctx.lineWidth = 1;
  ctx.setLineDash([4, 4]);

  const maxDist = 1200; // km
  const maxAlt = 300; // km

  // Vertical Grid Lines (Dist)
  for (let d = 200; d <= maxDist; d += 200) {
    const x = (d / maxDist) * (w - 40) + 30;
    ctx.beginPath();
    ctx.moveTo(x, 10);
    ctx.lineTo(x, h - 20);
    ctx.stroke();
  }

  // Horizontal Grid Lines (Alt)
  for (let a = 50; a <= maxAlt; a += 50) {
    const y = h - 20 - (a / maxAlt) * (h - 30);
    ctx.beginPath();
    ctx.moveTo(30, y);
    ctx.lineTo(w - 10, y);
    ctx.stroke();
  }

  // Draw Grid Axis Labels
  ctx.fillStyle = 'rgba(255, 255, 255, 0.3)';
  ctx.font = '9px monospace';
  ctx.textAlign = 'right';
  ctx.textBaseline = 'middle';
  ctx.setLineDash([]); // Reset line dash

  // Altitude Labels (Y-axis)
  for (let a = 0; a <= maxAlt; a += 100) {
    const y = h - 20 - (a / maxAlt) * (h - 30);
    ctx.fillText(`${a}`, 25, y);
  }

  // Downrange Labels (X-axis)
  ctx.textAlign = 'center';
  for (let d = 0; d <= maxDist; d += 400) {
    const x = (d / maxDist) * (w - 40) + 30;
    ctx.fillText(`${d}`, x, h - 8);
  }

  // Draw nominal trajectory path
  ctx.strokeStyle = 'rgba(6, 182, 212, 0.15)';
  ctx.lineWidth = 2;
  ctx.setLineDash([6, 3]);
  ctx.beginPath();
  
  const nominalPoints = [
    { d: 0, a: 0 },
    { d: 50, a: 25 },
    { d: 150, a: 80 },
    { d: 400, a: 180 },
    { d: 800, a: 240 },
    { d: 1200, a: 250 }
  ];

  nominalPoints.forEach((pt, idx) => {
    const x = (pt.d / maxDist) * (w - 40) + 30;
    const y = h - 20 - (pt.a / maxAlt) * (h - 30);
    if (idx === 0) ctx.moveTo(x, y);
    else ctx.lineTo(x, y);
  });
  ctx.stroke();
  ctx.setLineDash([]); // Reset

  // Draw actual trajectory path
  if (telemetryChartPoints.length > 0) {
    ctx.strokeStyle = telemetryAborted ? 'rgba(239, 68, 68, 0.8)' : 'rgba(6, 182, 212, 0.85)';
    ctx.lineWidth = 3;
    ctx.shadowBlur = 8;
    ctx.shadowColor = telemetryAborted ? 'rgba(239, 68, 68, 0.5)' : 'rgba(6, 182, 212, 0.5)';
    ctx.beginPath();

    telemetryChartPoints.forEach((pt, idx) => {
      const clampedX = Math.min(pt.x, maxDist);
      const clampedY = Math.min(pt.y, maxAlt);

      const x = (clampedX / maxDist) * (w - 40) + 30;
      const y = h - 20 - (clampedY / maxAlt) * (h - 30);
      if (idx === 0) ctx.moveTo(x, y);
      else ctx.lineTo(x, y);
    });
    ctx.stroke();
    
    ctx.shadowBlur = 0;

    const lastPt = telemetryChartPoints[telemetryChartPoints.length - 1];
    const clampedX = Math.min(lastPt.x, maxDist);
    const clampedY = Math.min(lastPt.y, maxAlt);
    const rX = (clampedX / maxDist) * (w - 40) + 30;
    const rY = h - 20 - (clampedY / maxAlt) * (h - 30);

    ctx.fillStyle = telemetryAborted ? '#ef4444' : '#06b6d4';
    ctx.beginPath();
    ctx.arc(rX, rY, 5, 0, Math.PI * 2);
    ctx.fill();

    const pulseRadius = 5 + (Math.sin(Date.now() / 150) + 1) * 3;
    ctx.strokeStyle = telemetryAborted ? 'rgba(239, 68, 68, 0.4)' : 'rgba(6, 182, 212, 0.4)';
    ctx.lineWidth = 1.5;
    ctx.beginPath();
    ctx.arc(rX, rY, pulseRadius, 0, Math.PI * 2);
    ctx.stroke();
  }

  ctx.restore();
}

function startTelemetrySimulation() {
  // Reset all state
  telemetryTime = -10;
  telemetrySpeed = 0;
  telemetryAlt = 0;
  telemetryDist = 0;
  telemetryStage1Fuel = 100;
  telemetryStage2Fuel = 100;
  telemetryAborted = false;
  telemetryThrottle = 100;
  telemetryChartPoints = [];

  // Setup UI references
  const speedEl = document.getElementById('telemetry-speed');
  const altEl = document.getElementById('telemetry-alt');
  const distEl = document.getElementById('telemetry-dist');
  
  if (speedEl) speedEl.innerHTML = `0 <span class="telemetry-unit-spacex">km/h</span>`;
  if (altEl) altEl.innerHTML = `0.0 <span class="telemetry-unit-spacex">km</span>`;
  if (distEl) distEl.innerHTML = `0.0 <span class="telemetry-unit-spacex">km</span>`;

  // Attach controls
  const throttleSlider = document.getElementById('throttle-override') as HTMLInputElement;
  const throttleVal = document.getElementById('throttle-val');
  if (throttleSlider) {
    throttleSlider.value = '100';
    throttleSlider.disabled = false;
    throttleSlider.addEventListener('input', (e) => {
      if (telemetryAborted) return;
      telemetryThrottle = parseInt((e.target as HTMLInputElement).value);
      if (throttleVal) throttleVal.textContent = `${telemetryThrottle}%`;
    });
  }

  const abortBtn = document.getElementById('abort-btn');
  if (abortBtn) {
    abortBtn.textContent = 'Abort Mission';
    abortBtn.className = 'abort-btn-spacex';
    abortBtn.addEventListener('click', () => {
      if (telemetryAborted) return;
      initiateAbort();
    });
  }

  const abortOverlay = document.getElementById('abort-overlay');
  if (abortOverlay) abortOverlay.classList.remove('active');

  // Trigger initial chart draw
  drawTelemetryChart();

  if (telemetryInterval) clearInterval(telemetryInterval);
  
  telemetryInterval = setInterval(() => {
    if (telemetryTime < 0) {
      telemetryTime += 0.1;
      const countdownTimeEl = document.getElementById('step-countdown-time');
      if (countdownTimeEl) countdownTimeEl.textContent = `T${telemetryTime.toFixed(1)}s`;
      
      const clockEl = document.getElementById('telemetry-clock');
      if (clockEl) clockEl.textContent = `T${telemetryTime.toFixed(1)}s`;
      
      if (Math.abs(telemetryTime) < 0.05) {
        telemetryTime = 0;
        addLogLine('LIFTOFF! Go Falcon Heavy! Go Psyche!', 'green');
        setActiveStep('step-liftoff');
      }
    } else {
      const dt = 2.0;
      telemetryTime += dt;

      const clockEl = document.getElementById('telemetry-clock');
      if (clockEl) {
        const minutes = Math.floor(telemetryTime / 60);
        const seconds = Math.floor(telemetryTime % 60);
        clockEl.textContent = `T+${minutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}`;
      }

      if (!telemetryAborted) {
        if (telemetryTime < 150) {
          telemetryStage1Fuel = Math.max(0, telemetryStage1Fuel - (telemetryThrottle / 100) * 1.33);
          
          telemetrySpeed += Math.round((telemetryThrottle / 100) * 40 * (1.2 - telemetryAlt / 100)); 
          telemetryAlt += (telemetryThrottle / 100) * 0.5 * (1 + telemetrySpeed / 4000);
          telemetryDist += (telemetrySpeed / 3600) * dt;

          if (telemetryTime >= 60 && telemetryTime - dt < 60) {
            setActiveStep('step-maxq');
            addLogLine('MAXIMUM AERODYNAMIC PRESSURE (MAX Q) REACHED.', 'cyan');
          }
        } else if (telemetryTime >= 150 && telemetryTime < 160) {
          telemetrySpeed = Math.max(0, telemetrySpeed - 5);
          telemetryAlt += 0.1;
          telemetryDist += (telemetrySpeed / 3600) * dt;

          if (telemetryTime >= 150 && telemetryTime - dt < 150) {
            setActiveStep('step-meco');
            addLogLine('MAIN ENGINE CUTOFF (MECO) CONFIRMED.', 'cyan');
          }
          if (telemetryTime >= 158 && telemetryTime - dt < 158) {
            setActiveStep('step-sep');
            addLogLine('STAGE 1 SEPARATION CONFIRMED.', 'green');
            addLogLine('[STAGE 2] Vacuum engine ignition.', 'cyan');
          }
        } else if (telemetryTime >= 160 && telemetryTime < 500) {
          telemetryStage2Fuel = Math.max(0, telemetryStage2Fuel - (telemetryThrottle / 100) * 0.29);

          telemetrySpeed += Math.round((telemetryThrottle / 100) * 62);
          if (telemetrySpeed > 27000) telemetrySpeed = 27000;

          telemetryAlt += 0.5 * (1 - (telemetryAlt - 80) / 250);
          if (telemetryAlt > 250) telemetryAlt = 250;

          telemetryDist += (telemetrySpeed / 3600) * dt;
        } else {
          telemetrySpeed = 27000;
          telemetryAlt = 250;
          telemetryDist += (telemetrySpeed / 3600) * dt;

          if (telemetryTime >= 500 && telemetryTime - dt < 500) {
            setActiveStep('step-orbit');
            addLogLine('ORBIT INSERTION COMPLETED. Nominal orbit established.', 'green');
            addLogLine('[MISSION SUCCESS] Flight controls normal.', 'green');
            clearInterval(telemetryInterval);
            telemetryInterval = null;
          }
        }
      } else {
        telemetrySpeed = Math.max(0, telemetrySpeed - 150);
        telemetryAlt = Math.max(0, telemetryAlt - 2.5);
        if (telemetryAlt > 0) {
          telemetryDist += (telemetrySpeed / 3600) * dt;
        } else {
          telemetrySpeed = 0;
        }

        if (telemetryAlt <= 0.05 && telemetryAlt > 0) {
          telemetryAlt = 0;
          addLogLine('[SAFETY] Ocean splashdown confirmed. Capsule telemetry terminated.', 'green');
          clearInterval(telemetryInterval);
          telemetryInterval = null;
        }
      }

      telemetryChartPoints.push({ x: telemetryDist, y: telemetryAlt });

      if (speedEl) speedEl.innerHTML = `${telemetrySpeed.toLocaleString()} <span class="telemetry-unit-spacex">km/h</span>`;
      if (altEl) altEl.innerHTML = `${telemetryAlt.toFixed(1)} <span class="telemetry-unit-spacex">km</span>`;
      if (distEl) distEl.innerHTML = `${telemetryDist.toFixed(1)} <span class="telemetry-unit-spacex">km</span>`;

      const fuel1Val = document.getElementById('fuel-stage1-val');
      const fuel1Fill = document.getElementById('fuel-stage1-fill');
      const fuel2Val = document.getElementById('fuel-stage2-val');
      const fuel2Fill = document.getElementById('fuel-stage2-fill');

      if (fuel1Val) fuel1Val.textContent = `${Math.round(telemetryStage1Fuel)}%`;
      if (fuel1Fill) fuel1Fill.style.width = `${telemetryStage1Fuel}%`;
      if (fuel2Val) fuel2Val.textContent = `${Math.round(telemetryStage2Fuel)}%`;
      if (fuel2Fill) fuel2Fill.style.width = `${telemetryStage2Fuel}%`;
    }

    drawTelemetryChart();
  }, 100);
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
function init() {
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

  // Setup Mobile Drawer Toggle
  const toggleBtn = document.getElementById('mobile-menu-toggle');
  const sidebar = document.querySelector('.sidebar');
  let overlay = document.querySelector('.sidebar-overlay');
  
  if (!overlay) {
    overlay = document.createElement('div');
    overlay.className = 'sidebar-overlay';
    document.body.appendChild(overlay);
  }

  if (toggleBtn && sidebar && overlay) {
    toggleBtn.addEventListener('click', () => {
      sidebar.classList.toggle('open');
      overlay.classList.toggle('visible');
    });

    overlay.addEventListener('click', () => {
      sidebar.classList.remove('open');
      overlay.classList.remove('visible');
    });

    // Auto-close sidebar on nav item click on mobile
    const navLinks = sidebar.querySelectorAll('.nav-menu .nav-item');
    navLinks.forEach(link => {
      link.addEventListener('click', () => {
        sidebar.classList.remove('open');
        overlay!.classList.remove('visible');
      });
    });
  }

  // Initial fetch for the default tab
  fetchData('launches');
}

if (document.readyState === 'loading') {
  document.addEventListener('DOMContentLoaded', init);
} else {
  init();
}

