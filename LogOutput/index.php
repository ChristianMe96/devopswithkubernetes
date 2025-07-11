<?php

use Random\RandomException;

// Generate random string on startup and store in memory
class LogOutputApp {
    private static $randomString;
    
    public static function init() {
        if (self::$randomString === null) {
            try {
                self::$randomString = bin2hex(random_bytes(18));
                error_log("Generated random string: " . self::$randomString);
            } catch (RandomException $e) {
                error_log("Failed to generate random string: " . $e->getMessage());
                self::$randomString = "error-generating-random-string";
            }
        }
    }
    
    public static function getRandomString() {
        return self::$randomString;
    }
}

// Initialize the application
LogOutputApp::init();

// Handle requests
$requestPath = $_SERVER['REQUEST_URI'] ?? '/';
$requestPath = parse_url($requestPath, PHP_URL_PATH);

if ($requestPath === '/status') {
    // Return JSON status
    header('Content-Type: application/json');
    echo json_encode([
        'timestamp' => date('Y-m-d H:i:s'),
        'random_string' => LogOutputApp::getRandomString()
    ]);
    exit;
}

// Default home page
?>
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Log Output App</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            max-width: 800px;
            margin: 0 auto;
            padding: 20px;
            background-color: #f5f5f5;
        }
        .container {
            background-color: white;
            padding: 30px;
            border-radius: 8px;
            box-shadow: 0 2px 4px rgba(0,0,0,0.1);
        }
        h1 {
            color: #333;
            border-bottom: 2px solid #007bff;
            padding-bottom: 10px;
        }
        .status-box {
            background-color: #f8f9fa;
            border: 1px solid #dee2e6;
            border-radius: 4px;
            padding: 15px;
            margin: 20px 0;
        }
        .endpoint-info {
            background-color: #d4edda;
            border: 1px solid #c3e6cb;
            border-radius: 4px;
            padding: 15px;
            margin: 20px 0;
        }
        code {
            background-color: #f8f9fa;
            padding: 2px 4px;
            border-radius: 3px;
            font-family: monospace;
        }
        a {
            color: #007bff;
            text-decoration: none;
        }
        a:hover {
            text-decoration: underline;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Log Output Application</h1>
        
        <div class="status-box">
            <h3>Current Status</h3>
            <p><strong>Random String:</strong> <code><?php echo htmlspecialchars(LogOutputApp::getRandomString()); ?></code></p>
            <p><strong>Generated at:</strong> <?php echo date('Y-m-d H:i:s'); ?></p>
        </div>
        
        <div class="endpoint-info">
            <h3>API Endpoint</h3>
            <p>Get current status as JSON: <a href="/status"><code>/status</code></a></p>
        </div>
        
        <p>This application generates a random string on startup and serves it via a web endpoint. The random string is stored in memory and persists as long as the application is running.</p>
    </div>
</body>
</html> 