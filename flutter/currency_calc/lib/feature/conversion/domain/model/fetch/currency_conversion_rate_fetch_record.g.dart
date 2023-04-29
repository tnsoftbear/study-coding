// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'currency_conversion_rate_fetch_record.dart';

// **************************************************************************
// TypeAdapterGenerator
// **************************************************************************

class CurrencyConversionRateFetchRecordAdapter
    extends TypeAdapter<CurrencyConversionRateFetchRecord> {
  @override
  final int typeId = 1;

  @override
  CurrencyConversionRateFetchRecord read(BinaryReader reader) {
    final numOfFields = reader.readByte();
    final fields = <int, dynamic>{
      for (int i = 0; i < numOfFields; i++) reader.readByte(): reader.read(),
    };
    return CurrencyConversionRateFetchRecord(
      sourceCurrency: fields[0] as String,
      targetCurrency: fields[1] as String,
      exchangeRate: fields[2] as double,
      createdAt: fields[3] as DateTime,
    );
  }

  @override
  void write(BinaryWriter writer, CurrencyConversionRateFetchRecord obj) {
    writer
      ..writeByte(4)
      ..writeByte(0)
      ..write(obj.sourceCurrency)
      ..writeByte(1)
      ..write(obj.targetCurrency)
      ..writeByte(2)
      ..write(obj.exchangeRate)
      ..writeByte(3)
      ..write(obj.createdAt);
  }

  @override
  int get hashCode => typeId.hashCode;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is CurrencyConversionRateFetchRecordAdapter &&
          runtimeType == other.runtimeType &&
          typeId == other.typeId;
}
