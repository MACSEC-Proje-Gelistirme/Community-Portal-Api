#!/bin/bash

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' 

echo "🔄 PostgreSQL veritabanı kurulumu başlatılıyor..."

if ! pg_isready > /dev/null 2>&1; then
    echo -e "${RED}❌ PostgreSQL servisi çalışmıyor!${NC}"
    echo "PostgreSQL servisinin çalıştığından emin olun."
    exit 1
fi

echo "📦 Veritabanı oluşturuluyor..."
if psql -U community -d community -f dbsetup.sql; then
    echo -e "${GREEN}✅ Veritabanı başarıyla kuruldu!${NC}"
else
    echo -e "${RED}❌ Veritabanı kurulumu sırasında bir hata oluştu!${NC}"
    exit 1
fi

if psql -U community -d community -c "SELECT 1" > /dev/null 2>&1; then
    echo -e "${GREEN}✅ Veritabanı bağlantısı başarılı!${NC}"
else
    echo -e "${RED}❌ Veritabanı bağlantısı başarısız!${NC}"
    exit 1
fi

echo -e "${GREEN}🎉 Kurulum tamamlandı!${NC}" 