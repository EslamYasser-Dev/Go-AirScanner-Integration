#!/bin/bash

# build_twain_dsm.sh
# Script to build TWAIN-DSM from source and prepare for Go cgo integration
# Usage: ./build_twain_dsm.sh

set -e  # Exit on any error

echo "🚀 Building TWAIN-DSM for Go integration..."

# Directories
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BUILD_ROOT="$SCRIPT_DIR/twain-dsm-build"
TWAIN_REPO="https://github.com/twain/twain-dsm.git"
TWAIN_DIR="$BUILD_ROOT/twain-dsm"
BUILD_DIR="$TWAIN_DIR/build"
DIST_DIR="$SCRIPT_DIR/dist"
INCLUDE_DIR="$DIST_DIR/include"
LIB_DIR="$DIST_DIR/lib"

# Create workspace
echo "📁 Creating build directory..."
mkdir -p "$BUILD_ROOT"
cd "$BUILD_ROOT"

# Clone TWAIN-DSM if not already cloned
if [ ! -d "$TWAIN_DIR" ]; then
    echo "📥 Cloning TWAIN-DSM from $TWAIN_REPO..."
    git clone "$TWAIN_REPO" "$TWAIN_DIR"
else
    echo "🔁 Using existing TWAIN-DSM repo."
    cd "$TWAIN_DIR"
    git pull origin main || git pull origin master || true
    cd "$BUILD_ROOT"
fi

# Create build directory
echo "⚙️ Configuring CMake..."
mkdir -p "$BUILD_DIR"
cd "$BUILD_DIR"

# Run CMake (Unix Makefiles by default)
cmake .. -DCMAKE_BUILD_TYPE=Release

# Build
echo "🔨 Building TWAIN-DSM..."
cmake --build . --config Release --target ALL_BUILD

# Create dist structure
echo "📦 Creating distribution directories..."
mkdir -p "$INCLUDE_DIR"
mkdir -p "$LIB_DIR"

# Copy headers
echo "📑 Copying headers to $INCLUDE_DIR..."
cp -r "$TWAIN_DIR/Include/"*.h "$INCLUDE_DIR/"

# Copy library (Linux/macOS/Windows)
if [ -f "libTWAINDSM.a" ]; then
    # Linux/macOS static lib
    cp "libTWAINDSM.a" "$LIB_DIR/"
elif [ -f "libTWAINDSM.so" ]; then
    # Shared lib (Linux)
    cp "libTWAINDSM.so" "$LIB_DIR/"
elif [ -f "TWAINDSM.dll" ]; then
    # Windows DLL
    cp "TWAINDSM.dll" "$LIB_DIR/"
    cp "TWAINDSM.lib" "$LIB_DIR/" 2>/dev/null || echo "⚠️ TWAINDSM.lib not found (optional)"
elif [ -f "Release/TWAINDSM.lib" ]; then
    # Windows: Visual Studio places in Release/
    cp "Release/TWAINDSM.lib" "$LIB_DIR/"
    cp "Release/TWAINDSM.dll" "$LIB_DIR/"
fi

# Success
echo ""
echo "✅ TWAIN-DSM built successfully!"
echo "📁 Headers: $INCLUDE_DIR"
echo "📁 Libraries: $LIB_DIR"
echo ""
echo "🔗 Now use in Go with:"
echo '   #cgo CXXFLAGS: -I./dist/include'
echo '   #cgo LDFLAGS: -L./dist/lib -lTWAINDSM -lole32 -loleaut32 -luser32'
echo ""