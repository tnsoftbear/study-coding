// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'currency_conversion_history_record.dart';

// **************************************************************************
// TypeAdapterGenerator
// **************************************************************************

class CurrencyConversionHistoryRecordAdapter
    extends TypeAdapter<CurrencyConversionHistoryRecord> {
  @override
  final int typeId = 0;

  @override
  CurrencyConversionHistoryRecord read(BinaryReader reader) {
    final numOfFields = reader.readByte();
    final fields = <int, dynamic>{
      for (int i = 0; i < numOfFields; i++) reader.readByte(): reader.read(),
    };
    return CurrencyConversionHistoryRecord()
      ..sourceCurrency = fields[0] as String
      ..sourceAmount = fields[1] as double
      ..targetCurrency = fields[2] as String
      ..targetAmount = fields[3] as double
      ..rate = fields[4] as double
      ..date = fields[5] as DateTime?;
  }

  @override
  void write(BinaryWriter writer, CurrencyConversionHistoryRecord obj) {
    writer
      ..writeByte(6)
      ..writeByte(0)
      ..write(obj.sourceCurrency)
      ..writeByte(1)
      ..write(obj.sourceAmount)
      ..writeByte(2)
      ..write(obj.targetCurrency)
      ..writeByte(3)
      ..write(obj.targetAmount)
      ..writeByte(4)
      ..write(obj.rate)
      ..writeByte(5)
      ..write(obj.date);
  }

  @override
  int get hashCode => typeId.hashCode;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is CurrencyConversionHistoryRecordAdapter &&
          runtimeType == other.runtimeType &&
          typeId == other.typeId;
}